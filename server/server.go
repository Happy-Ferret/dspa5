package main

import "os"
import "fmt"
import "regexp"
import "golang.org/x/net/context"
import "google.golang.org/grpc"
import pb "github.com/naggie/dspa5/dspa5"

const port = ":40401"

type fragment struct {
	text        string
	wavFilepath string
}

type server struct {
	synthQueue chan fragment
	playQueue  chan fragment
	doneQueue  chan fragment
}

func NewServer() *server {
	return &server{
		make(chan fragment, 10),
		make(chan fragment, 10),
		make(chan fragment, 10),
	}
}

func (s *server) Speak(annoucement *pb.Announcement, stream pb.Dspa5_SpeakServer) error {
	texts := regexp.MustCompile(": |;|,|\\.|(?<=\\!) |(?<=\\?) ").Split(annoucement.Message, -1)
	for text := range texts {
		
	}

	for f := range s.doneQueue {
		err := stream.Send(&pb.Announcement{f.text, pb.Announcement_WARNING})
	}
}

func main() {
	path, ok := os.LookupEnv("DSPA_TTS_COMMAND")

	if !ok {
		panic("DSPA_TTS_COMMAND not set")
	}

	s := NewServer()
	go s.synthWorker()
	go s.playWorker()

	fmt.Printf("Path is %v\n", path)
}

func (s *server) synthWorker() {
	for f := range s.synthQueue {
		if f.text != "" {
			f.wavFilepath = synth(f.text)
		}

		s.playQueue <- f
	}
}

func (s *server) playWorker() {
	for f := range s.playQueue {
		play(f.wavFilepath)
		s.doneQueue <- f
	}
}

func synth(text string) string {

}

func play(filepath string) {

}
