package main

import (
	pb "github.com/naggie/dspa5/dspa5"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	path, ok := os.LookupEnv("DSPA_SERVER_ADDR")

	if !ok {
		log.Fatalf("DSPA_TTS_COMMAND not set")
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewDspa5Client(conn)
}
