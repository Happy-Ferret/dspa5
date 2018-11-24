package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title: "DSPA Display",
		Monitor: pixelgl.PrimaryMonitor(),
		Undecorated: true,
		VSync: true,
	}
}

func main() {
	pixelgl.Run(run)
}
