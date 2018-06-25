package main

import (
	"fmt"
	pb "github.com/naggie/dspa5/dspa5"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/exec"
	"sync"
	"strings"
	"time"
	"ioutil"
	"crypto/sha256"
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
	f, err := TempFile(dir, "synth")

	// before https://go-review.googlesource.com/c/go/+/105675
	name := f.Name() + '.aiff'
	os.Rename(f.Name(), newpath string)
	f, err := os.Create(name)


	tmpfile := tmpfile
	cmd := exec.Command("say", "-o", tmpfile)


	cmd.Run()

	if err != nil {
		log.Fatalf("Error running synth: %v\n", err)
	}

	return filepath
}

func play(filepath string) {
	err := exec.Command("play", filepath).Run()

	if err != nil {
		log.Fatalf("Error running play: %v\n", err)
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

func requireEnv(key string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		log.Fatalf("%v required in environment", key)
	}

	return val
}

func main() {
	tmpDir = requireEnv("DSPA_DATA_DIR") TODO create subdirs!
	dataDir = requireEnv("DSPA_DATA_DIR")
	synthCmd = requireEnv("DSPA_SYNTH_CMD")
	playCmd = requireEnv("DSPA_PLAY_CMD")

	lis, err := net.Listen("tcp", "0.0.0.0:55223")

	if err == nil {
		log.Infof("Listening on port 55223")
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
