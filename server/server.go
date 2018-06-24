package main

import "os"
import "fmt"
import "sync"
import "net"
import "log"
import "regexp"
import "google.golang.org/grpc"
import pb "github.com/naggie/dspa5/dspa5"

const port = ":40401"

var startChimes = map[int]string{
	pb.Announcement_INFO:     "xerxes_start.wav",
	pb.Announcement_WARNING:  "warning.wav",
	pb.Announcement_ERROR:    "error.wav",
	pb.Announcement_CRITICAL: "redalert.wav",
}

var stopChimes = map[int]string{
	pb.Announcement_INFO:     "xerxes_stop.wav",
	pb.Announcement_CRITICAL: "redalert.wav",
}

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
	s.announcementLock.Lock()

	playingChannel := make(chan *fragment, 10)

	s.synthQueue <- &fragment{"", startChimes[announcement.Level], playingChannel, false}

	texts := regexp.MustCompile(": |;|,|\\.|(?<=\\!) |(?<=\\?) ").Split(announcement.Message, -1)
	for _, text := range texts {
		s.synthQueue <- &fragment{text, "", playingChannel, false}
	}

	// send combined chime + stop marker to close channel on completion
	s.synthQueue <- &fragment{"", startChimes[announcement.Level], playingChannel, true}

	s.announcementLock.Unlock()

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
	return "a filepath"
}

func play(filepath string) {
	fmt.Println(filepath)
	return
}

func main() {
	path, ok := os.LookupEnv("DSPA_TTS_COMMAND")

	if !ok {
		log.Fatalf("DSPA_TTS_COMMAND not set")
	}

	lis, err := net.Listen("tcp", "localhost:55223")

	if err != nil {
		log.Fatalf("Failed to listen on port 55223")
	}

	s := NewServer()
	go s.synthWorker()
	go s.playWorker()

	fmt.Printf("Path is %v\n", path)

	grpcServer := grpc.NewServer()
	pb.RegisterDspa5Server(grpcServer, s)
	grpcServer.Serve(lis)
}
