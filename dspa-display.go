package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	monitor := pixelgl.PrimaryMonitor()
	width, height := monitor.Size()

	cfg := pixelgl.WindowConfig{
		Title: "DSPA Display",
		Monitor: monitor,
		Undecorated: true,
		VSync: true,
		Bounds: pixel.R(0, 0, width, height),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Update()
	}
}

// GUI thread must be on main thread for most OSes which is difficult in go.
// This achieves it.
func main() {
	pixelgl.Run(run)
}
