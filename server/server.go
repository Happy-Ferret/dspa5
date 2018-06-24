package main

import "os"
import "fmt"
import "regexp"
import "golang.org/x/net/context"
import "google.golang.org/grpc"
import pb "github.com/naggie/dspa5/dspa5"

const port = ":40401"


const (
	NOTSET = 0
	DEBUG = 10
	INFO = 20
	WARNING = 30
	ERROR = 40
	CRITICAL = 50
)

startChimes =: map[int]string{
	INFO : "xerxes_start.wav",
	WARNING : "warning.wav",
	ERROR : "error.wav",
	CRITICAL : "redalert.wav",
}

stopChimes =: map[int]string{
	INFO : "xerxes_stop.wav",
	CRITICAL : "redalert.wav",
}

type fragment struct {
	// optional text to say
	text        string
	// optional wav file to play (added by synth if necessary)
	wavFilepath string
	// report chimes/speech as it happens (TODO must be per annoucement)
	playingChannel chan *fragment
	// is this the last message for the request associated with playingChannel? If it
	// is, will be closed after play.
	last bool
}

type server struct {
	// used to serialise announcements
	announcementQueue chan *pb.Announcement
	// synthesise speech if any (or pass on chime)
	synthQueue chan *fragment
	// play chime or speech
	playQueue  chan *fragment
}

func NewServer() *server {
	return &server{
		make(chan *pb.Announcement, 10),
		make(chan *fragment, 10),
		make(chan *fragment, 10),
	}
}

func (s *server) Speak(annoucement *pb.Announcement, stream pb.Dspa5_SpeakServer) error {

	playingChannel := make(chan *fragment, 10)

	texts := regexp.MustCompile(": |;|,|\\.|(?<=\\!) |(?<=\\?) ").Split(annoucement.Message, -1)
	for _, text := range texts {
		s.synthQueue <- &fragment{text, "", playingChannel, false}
	}

	// send stop marker to close channel on completion
	s.synthQueue <- &fragment{"", "", playingChannel, true}

	for f := range doneQueue {
		err := stream.Send(&pb.Announcement{f.text, announcement.Level})
	}
}

func main() {
	path, ok := os.LookupEnv("DSPA_TTS_COMMAND")

	if !ok {
		panic("DSPA_TTS_COMMAND not set")
	}

	s := NewServer()
	go s.synthWorker()
	go s.playWorker()

	fmt.Printf("Path is %v\n", path)
}

func (s *server) synthWorker() {
	for f := range s.synthQueue {
		if f.text != "" {
			f.wavFilepath = synth(f.text)
		}

		s.playQueue <- f
	}
}

func (s *server) playWorker() {
	for f := range s.playQueue {
		play(f.wavFilepath)
		s.doneQueue <- f
	}
}

func synth(text string) string {

}

func play(filepath string) {

}
