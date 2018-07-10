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
}
