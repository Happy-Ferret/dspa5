package main

import "os"
import "io"
import "fmt"
import "golang.org/x/net/context"
import "google.golang.org/grpc"
import pb "github.com/naggie/dspa5/proto"

const port = ":40401"

type fragment struct {
	text         string
	wav_filepath string
}

type server struct {
	synth_queue chan fragment
	play_queue  chan fragment
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

	fmt.Printf("Path is %v\n", path)
}
