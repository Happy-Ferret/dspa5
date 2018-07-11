package main

import (
	"fmt"
	pb "github.com/naggie/dspa5/dspa5"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "\nUsage: dspa-client [level] message...\n\n")
		fmt.Fprintf(os.Stderr, "Where optional level is one of INFO, WARNING, ERROR, CRITICAL\n")
		fmt.Fprintf(os.Stderr, "Default level is INFO\n\n")
		os.Exit(1)
	}

	// conveniently, will be 0 = NOTSET for any invalid value
	level := pb.Announcement_Level(pb.Announcement_Level_value[os.Args[1]])

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

	var message string

	if level == pb.Announcement_NOTSET {
		message = strings.Join(os.Args[1:], " ")
	} else {
		// ignore level from args
		message = strings.Join(os.Args[2:], " ")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	stream, err := client.Speak(ctx, &pb.Announcement{Message: message, Level: level})

	if err != nil {
		log.Fatalf("Failed to announce: %v", err)
	}

	for {
		fragment, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to stream: %v", err)
		}

		if fragment.Error {
			fmt.Println("[ERROR]")
		} else if fragment.Chime {
			fmt.Println("[CHIME]")
		} else {
			fmt.Println(fragment.Text)
		}
	}
}
