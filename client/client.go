package main

import (
	pb "github.com/naggie/dspa5/dspa5"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
	"golang.org/x/net/context"
	"time"
	"io"
)

func main() {
	serverAddr, ok := os.LookupEnv("DSPA_SERVER_ADDR")

	if !ok {
		log.Fatalf("DSPA_SERVER_ADDR not set")
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewDspa5Client(conn)

	message := strings.Join(os.Args[1:], " ")

	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()

	stream, err := client.Speak(ctx, &pb.Announcement{message, pb.Announcement_WARNING})

	if err != nil {
		log.Fatalf("Failed to announce: %v", err)
	}

	for {
		announcement, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to stream: %v", err)
		}

		log.Println(announcement.Message)
	}
}
