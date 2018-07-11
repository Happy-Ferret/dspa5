package main

import (
	pb "github.com/naggie/dspa5/dspa5"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"strings"
	"time"
	sd01 "github.com/naggie/sd01/go"
)

func main() {
	discoverer := sd01.NewDiscoverer('dspa5speaker')
	discoverer.Start()

	server := NewServer(discoverer)
}

type server struct {
	// used to serialise announcements. Lock is used as insertion is
	// nonblocking, easier than using channel.
	announcementLock *sync.Mutex
}

func NewServer(discoverer sd01.Discoverer) *server {
	return &server{
		discoverer: discoverer,
		announcementLock: &sync.Mutex{},
	}
}

func (s *server) Speak(announcement pb.Announcement, stream pb.Dspa5_SpeakServer) error {
	hosts := s.discoverer.GetServices()
	fragments := make(chan pb.Fragment, 10)

	for i, host := range hosts {
		// listen to first one only
		if i == 0 {
			speakUpstream(host, announcement, fragments)
		} else {
			speakUpstream(host, announcement, nil)
		}
	}

	for fragment := range fragments {
		stream.send(fragment)
	}
}

func speakUpstream(host string, announcement pb.Announcement, fragments chan<- pb.Fragment) error {
	if fragments != nil {
		defer close(fragments)
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)

	if err != nil {
		log.Printf("Failed to connect: %v", err)
		return err
	}

	defer conn.Close()

	client := pb.NewDspa5Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	stream, err := client.Speak(ctx, announcement)

	if err != nil {
		log.Printf("Failed to announce: %v", err)
		return err
	}

	for {
		fragment, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Failed to stream: %v", err)
			return err
		}

		if fragments != nil {
			fragments <- fragment
		}

	}

}
