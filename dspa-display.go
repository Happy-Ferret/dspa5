package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"io/ioutil"
	"os"
	"fmt"
	"image"
	_ "image/png"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"time"
)


// TODO define update once per change instead of regularly

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

	win.SetSmooth(true)
	win.Clear(colornames.Black)

	logo, err := loadPicture("logo.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(logo, logo.Bounds())
	sprite.Draw(win, pixel.IM.Moved(pixel.V(width/2, 2*height/3)))


	face, err := loadTTF("etc/raleway/Raleway-Regular.ttf", 80)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(face, text.ASCII)
	txt := text.New(pixel.V(width/2, height/3), atlas)
	txt.Color = colornames.White

	lines := []string {
		"This is a line wrapped multiline",
		"String. NEEDS UPPERCASE",
	}

	for _, line := range lines {
		txt.Dot.X -= txt.BoundsOf(line).W() / 2
		fmt.Fprintln(txt, line)
	}

	txt.Draw(win, pixel.IM)
	win.Update()
	win.Update()
	time.Sleep(time.Second)

	//for !win.Closed() {
	//	win.Update()
	//}
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

func loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}
