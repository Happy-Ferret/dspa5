package main

import "os"
import "fmt"
import "golang.org/x/net/context"
import "google.golang.org/grpc"
import pb "github.com/naggie/dspa5/proto"

const port ":40401"

type server struct{}


func (s *server) Speak(ctx context.Context, announcement *pb.Announcement) (*pb.Announcement, error) {

}



func main() {
	path, ok := os.LookupEnv("DSPA_TTS_COMMAND")

	if !ok {
		panic("DSPA_TTS_COMMAND not set")
	}

	fmt.Printf("Path is %v\n", path)
}
