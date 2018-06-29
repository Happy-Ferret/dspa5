package main

import (
	"crypto/sha256"
	"encoding/hex"
	pb "github.com/naggie/dspa5/dspa5"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
)

const port = ":40401"

var startChimes = map[pb.Announcement_Level]string{
	pb.Announcement_INFO:     "chimes/xerxes_start.ogg",
	pb.Announcement_WARNING:  "chimes/xerxes_chime.ogg",
	pb.Announcement_ERROR:    "chimes/xerxes_motion.wav",
	pb.Announcement_CRITICAL: "chimes/xerxes_motion.wav",
}

var stopChimes = map[pb.Announcement_Level]string{
	pb.Announcement_INFO:     "chimes/xerxes_stop.ogg",
	pb.Announcement_CRITICAL: "chimes/xerxes_breech.wav",
}

var tmpDir string
var cacheDir string
var synthCmd string
var playCmd []string
var fileExt string

type fragment struct {
	// optional text to say
	text string
	// optional wav file to play (added by synth if necessary)
	wavFile string
	// report chimes/speech as it happens
	playingChannel chan *fragment
	// is this the last message for the request associated with playingChannel? If it
	// is, will be closed after play. Can be an additional message with only this flag set.
	last bool
	// if an error occurred
	synthErr error
	playErr  error
}

type server struct {
	// used to serialise announcements. Lock is used as insertion is
	// nonblocking, easier than using channel.
	announcementLock *sync.Mutex
	// synthesise speech if any (or pass on chime)
	synthQueue chan *fragment
	// play chime or speech
	playQueue chan *fragment
}

func NewServer() *server {
	return &server{
		&sync.Mutex{},
		make(chan *fragment, 10),
		make(chan *fragment, 10),
	}
}

func (s *server) Speak(announcement *pb.Announcement, stream pb.Dspa5_SpeakServer) error {
	playingChannel := make(chan *fragment, 10)

	if announcement.Level == 0 {
		announcement.Level = pb.Announcement_INFO
	}

	// lock to serialise announcement messages so fragments don't interleave
	s.announcementLock.Lock()

	// send start chime
	s.synthQueue <- &fragment{"", startChimes[announcement.Level], playingChannel, false, nil, nil}

	// split message into text fragments to synthesise separately
	texts := split(announcement.Message)
	for _, text := range texts {
		s.synthQueue <- &fragment{text, "", playingChannel, false, nil, nil}
	}

	// send combined chime + stop marker to close channel on completion
	s.synthQueue <- &fragment{"", stopChimes[announcement.Level], playingChannel, true, nil, nil}

	s.announcementLock.Unlock()

	// read announcement fragments back to the client as they happen
	for f := range playingChannel {
		stream.Send(&pb.Announcement{f.text, announcement.Level})
	}

	return nil
}

func (s *server) synthWorker() {
	for f := range s.synthQueue {
		if f.text != "" {
			// error can be safely ignored -- synth logs and play won't play
			// nothing
			f.wavFile, f.synthErr = synth(f.text)
		}

		s.playQueue <- f
	}
}

func (s *server) playWorker() {
	for f := range s.playQueue {
		f.playingChannel <- f

		if f.wavFile != "" {
			f.playErr = play(f.wavFile)
		}

		if f.last {
			close(f.playingChannel)
		}
	}
}

func synth(text string) (string, error) {
	hash := sha256.Sum256([]byte(text))
	cacheFile := path.Join(cacheDir, hex.EncodeToString(hash[:])+"."+fileExt)

	if _, err := os.Stat(cacheFile); err == nil {
		return cacheFile, nil
	}

	f, err := ioutil.TempFile(tmpDir, "synth")
	if err != nil {
		log.Printf("Error creating tmpfile: %v", err)
		return "", err
	}

	// before https://go-review.googlesource.com/c/go/+/105675
	tmpFile := f.Name() + fileExt
	os.Rename(f.Name(), tmpFile)
	os.Remove(tmpFile)

	args := strings.Split(synthCmd + " " + tmpFile, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = strings.NewReader(text)

	if err = cmd.Run(); err != nil {
		log.Printf("Error running synth: %v", err)
		return "", err
	}

	// atomic!
	if err = os.Rename(tmpFile, cacheFile); err != nil {
		log.Printf("Error moving into tmp file: %v", err)
		return "", err
	}

	return cacheFile, nil
}

func play(filepath string) error {
	err := exec.Command("play", filepath).Run()

	if err != nil {
		log.Printf("Error running play: %v\n", err)
	}

	return err
}

func split(message string) []string {
	// replace punctuation with linebreaks, preserving ! and ?
	message = strings.Replace(message, ": ", "\n", -1)
	message = strings.Replace(message, "; ", "\n", -1)
	message = strings.Replace(message, ", ", "\n", -1)
	message = strings.Replace(message, ". ", "\n", -1)
	message = strings.Replace(message, "! ", "!\n", -1)
	message = strings.Replace(message, "? ", "?\n", -1)

	return strings.Split(message, "\n")
}

func requireEnv(key string, description string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		log.Fatalf("%v (%v) required in environment", key, description)
	}

	return val
}

func main() {
	tmpDir = path.Join(requireEnv("DSPA_DATA_DIR", "Directory to store tmp files and cache"), "tmp/")
	cacheDir = path.Join(requireEnv("DSPA_DATA_DIR", ""), "cache/")

	if err := os.MkdirAll(tmpDir, os.ModePerm); err != nil {
		log.Fatalf("Could not create %v: %v", tmpDir, err)
	}

	if err := os.MkdirAll(cacheDir, os.ModePerm); err != nil {
		log.Fatalf("Could not create %v: %v", cacheDir, err)
	}

	synthCmd = requireEnv("DSPA_SYNTH_CMD", "Command that accepts text on stdin and file to write on argv[1]")
	playCmd = strings.Split(requireEnv("DSPA_PLAY_CMD", "Command to play an audio file"), " ")

	fileExt = requireEnv("DSPA_FILE_EXT", "File extension of audio files with the dot")

	lis, err := net.Listen("tcp", "0.0.0.0:55223")

	if err == nil {
		log.Printf("Listening on port 55223")
	} else {
		log.Fatalf("Failed to listen on port 55223")
	}

	s := NewServer()
	go s.synthWorker()
	go s.playWorker()

	grpcServer := grpc.NewServer()
	pb.RegisterDspa5Server(grpcServer, s)
	grpcServer.Serve(lis)
}
