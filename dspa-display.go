package main

//go:generate go-bindata -pkg dspa5 -o dspa5/displayassets.go etc/roboto/Roboto-Regular.ttf

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
	"time"
	"strings"
)


func run() {
	monitor := pixelgl.PrimaryMonitor()

	logo := mustLoadPicture("logo.png")
	atlas := mustLoadTTFAtlas("etc/roboto/Roboto-Regular.ttf", 80)

	splash := NewSplash(monitor, logo, atlas)
	splash.SetLogo(true)
	splash.SetText("Hello world!")
	splash.DoFrame()
	splash.DoFrame()
	time.Sleep(2*time.Second)

	//for !win.Closed() {
	//	win.Update()
	//}
}

// GUI thread must be on main thread for most OSes which is difficult in go.
// This achieves it.
func main() {
	pixelgl.Run(run)
}

type Splash struct {
	monitor *pixelgl.Monitor
	window *pixelgl.Window
	logo pixel.Picture
	sprite *pixel.Sprite
	atlas *text.Atlas
	width float64
	height float64
	logoVisible bool
	text string
}

func NewSplash(monitor *pixelgl.Monitor, logo pixel.Picture, atlas *text.Atlas) *Splash {
	width, height := monitor.Size()
	s := Splash{
		monitor: monitor,
		logo: logo,
		atlas: atlas,
		width: width,
		height: height,
	}

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
	win.SetCursorVisible(false)
	s.window = win

	s.sprite = pixel.NewSprite(logo, logo.Bounds())

	return &s
}

func (s *Splash) SetLogo(visible bool) {
	s.logoVisible = visible
}

func (s *Splash) SetText(text string) {
	s.text = text
}

func (s *Splash) DoFrame() {
	s.window.Clear(colornames.Black)

	// default text position with no logo -- center (80px font)
	textY := s.height/2 - 40

	if s.logoVisible {
		s.sprite.Draw(s.window, pixel.IM.Moved(pixel.V(s.width/2, 2*s.height/3)))
		// to position the text half way between the logo and bottom (80px font size)
		textY = ((2*s.height/3 - s.logo.Bounds().H()/2))/2
	}

	txt := text.New(pixel.V(s.width/2, textY), s.atlas)
	txt.Color = colornames.White
	line := strings.ToUpper(s.text)
	txt.Dot.X -= txt.BoundsOf(line).W() / 2
	fmt.Fprintln(txt, line)
	txt.Draw(s.window, pixel.IM)

	s.window.Update()
}

func (s *Splash) Destroy() {
	s.window.Destroy()
}

func mustLoadPicture(path string) pixel.Picture {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	return pixel.PictureDataFromImage(img)
}

func mustLoadTTFAtlas(path string, size float64) *text.Atlas {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		panic(err)
	}

	face := truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	})

	return text.NewAtlas(face, text.ASCII)
}

