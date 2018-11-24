package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"os"
	"image"
	_ "image/png"
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

	win.Clear(colornames.Black)

	logo, err := loadPicture("logo.png")
	sprite := pixel.NewSprite(logo, logo.Bounds())
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

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

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
