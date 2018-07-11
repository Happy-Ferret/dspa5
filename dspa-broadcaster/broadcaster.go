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
	replies := make(chan pbAnnouncement, 10)

	for host := range hosts {
		speakUpstream(host, announcement, replies)
	}
}

func speakUpstream(host string, announcement pb.Announcement, replies chan<- pb.Announcement) {

}
