package main

import "os"
import "io"
import "fmt"
import "golang.org/x/net/context"
import "google.golang.org/grpc"
import pb "github.com/naggie/dspa5/proto"

const port ":40401"

type server struct{
	//synth chan
}


func (s *server) Speak(stream pb.Dspa5_SpeakServer) error {
	for {
		point, err := stream.Recv()

		if err == io.EOF {
			return ...something
		}

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
