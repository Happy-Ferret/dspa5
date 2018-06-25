package main

import (
	"crypto/sha256"
	pb "github.com/naggie/dspa5/dspa5"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
	"io"
	"encoding/hex"
	"io/ioutil"
	"log"
)

const port = ":40401"

var startChimes = map[pb.Announcement_Level]string{
	pb.Announcement_INFO:     "xerxes_start.wav",
	pb.Announcement_WARNING:  "warning.wav",
	pb.Announcement_ERROR:    "error.wav",
	pb.Announcement_CRITICAL: "redalert.wav",
}

var stopChimes = map[pb.Announcement_Level]string{
	pb.Announcement_INFO:     "xerxes_stop.wav",
	pb.Announcement_CRITICAL: "redalert.wav",
}

var tmpDir string
var cacheDir string
var synthCmd string
var playCmd string
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
	// lock to serialise announcement messages so fragments don't interleave
	s.announcementLock.Lock()

	playingChannel := make(chan *fragment, 10)

	// send start chime
	s.synthQueue <- &fragment{"", startChimes[announcement.Level], playingChannel, false}

	// split message into text fragments to synthesise separately
	texts := split(announcement.Message)
	for _, text := range texts {
		s.synthQueue <- &fragment{text, "", playingChannel, false}
	}

	// send combined chime + stop marker to close channel on completion
	s.synthQueue <- &fragment{"", startChimes[announcement.Level], playingChannel, true}

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
			f.wavFile = synth(f.text)
		}

		s.playQueue <- f
	}
}

func (s *server) playWorker() {
	for f := range s.playQueue {
		f.playingChannel <- f

		if f.wavFile != "" {
			play(f.wavFile)
		}

		if f.last {
			close(f.playingChannel)
		}
	}
}

func synth(text string) string {
	hash := sha256.Sum256([]byte(text))
	cacheFile := path.Join(cacheDir, hex.EncodeToString(hash) + fileExt)

	f, err := ioutil.TempFile(tmpDir, "synth")

	// before https://go-review.googlesource.com/c/go/+/105675
	name := f.Name() + fileExt
	os.Rename(f.Name(), name)
	defer os.Remove(name)

	cmd := exec.Command("say", "-o", name)

	stdin, err := cmd.StdinPipe()

	if err != nil {
		log.Println("Error opening stdin: %v", err)

	}
	defer stdin.Close()

	if err = cmd.Start(); err != nil {
		log.Println("Error starting synth: %v", err)
	}

	io.WriteString(stdin, text)

	if err = cmd.Wait(); err != nil {
		log.Println("Error running synth: %v", err)
	}

	return "filepath of object in cache"
}

func play(filepath string) {
	err := exec.Command("play", filepath).Run()

	if err != nil {
		log.Printf("Error running play: %v\n", err)
	}
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
	cacheDir = path.Join(requireEnv("DSPA_DATA_DIR",""), "cache/")
	os.MkdirAll(tmpDir, os.ModePerm)
	os.MkdirAll(cacheDir, os.ModePerm)

	synthCmd = requireEnv("DSPA_SYNTH_CMD", "Command that accepts text on stdin and file to write on argv[1]")
	playCmd = requireEnv("DSPA_PLAY_CMD", "Command to play an audio file")
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
