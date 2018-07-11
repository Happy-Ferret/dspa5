package main

import (
	pb "github.com/naggie/dspa5/dspa5"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
	sd01 "github.com/naggie/sd01/go"
	"sync"
	"net"
)

func main() {
	discoverer := sd01.NewDiscoverer("dspa5speaker")
	discoverer.Start()

	s := NewServer(discoverer)
	grpcServer := grpc.NewServer()
	pb.RegisterDspa5Server(grpcServer, s)
	lis, err := net.Listen("tcp", "0.0.0.0:55224")

	if err == nil {
		log.Printf("Listening on port 55224")
	} else {
		log.Fatalf("Failed to listen on port 55224")
	}

	grpcServer.Serve(lis)
}

type server struct {
	discoverer *sd01.Discoverer
	// used to serialise announcements. Lock is used as insertion is
	// nonblocking, easier than using channel.
	announcementLock *sync.Mutex
}

func NewServer(discoverer *sd01.Discoverer) *server {
	return &server{
		discoverer: discoverer,
		announcementLock: &sync.Mutex{},
	}
}

func (s *server) Speak(announcement *pb.Announcement, stream pb.Dspa5_SpeakServer) error {
	services := s.discoverer.GetServices()
	fragments := make(chan *pb.Fragment, 10)

	for i, service := range services {
		serverAddr := service.String()
		// listen to first one only
		if i == 0 {
			go speakUpstream(serverAddr, announcement, fragments)
		} else {
			go speakUpstream(serverAddr, announcement, nil)
		}
	}

	for fragment := range fragments {
		stream.Send(fragment)
	}

	return nil
}

func speakUpstream(serverAddr string, announcement *pb.Announcement, fragments chan<- *pb.Fragment) error {
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
	return nil
}
