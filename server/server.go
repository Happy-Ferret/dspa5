package main

import "os"
import "io"
import "fmt"
import "golang.org/x/net/context"
import "google.golang.org/grpc"
import pb "github.com/naggie/dspa5/proto"

const port = ":40401"

type fragment struct {
	text        string
	wavFilepath string
}

type server struct {
	synthQueue chan fragment
	playQueue  chan fragment
}

func NewServer() *server {
	return &server{
		make(chan fragment, 10),
		make(chan fragment, 10),
	}
}

func (s *server) Speak(annoucement *pb.Announcement, stream pb.Dspa5_SpeakServer) error {
	for {
		fmt.Printf("annoucement: %v\n", annoucement)

		err = stream.Send(&pb.Announcement{"Hello", pb.Announcement_WARNING})
	}
}

func main() {
	path, ok := os.LookupEnv("DSPA_TTS_COMMAND")

	if !ok {
		panic("DSPA_TTS_COMMAND not set")
	}

	s = NewServer()
	go s.synthWorker()
	go s.playWorker()

	fmt.Printf("Path is %v\n", path)
}

func (s *server) synthWorker() {
	for f := range s.synthQueue {
		if f.text {
			f.wavFilepath = synth(f.text)
		}

		playQueue <- f
	}
}

func (s *server) playWorker() {
	for f := range s.playQueue {
		play(f.wavFilepath)
	}
}

func getFragments(message string) string {

}

func synth(text string) string {

}

func play(filepath string) {

}
