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
	"strings"
)


func run() {
	monitor := pixelgl.PrimaryMonitor()

	logo, err := loadPicture("logo.png")
	if err != nil {
		panic(err)
	}

	face, err := loadTTF("etc/roboto/Roboto-Regular.ttf", 80)
	if err != nil {
		panic(err)
	}
	atlas := text.NewAtlas(face, text.ASCII)

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

	// default text position with no logo -- center
	textY := s.height/2 -40

	if s.logoVisible {
		s.sprite.Draw(s.window, pixel.IM.Moved(pixel.V(s.width/2, 2*s.height/3)))
		// to position the text half way between the logo and bottom (80px font size)
		textY = ((2*s.height/3 - s.logo.Bounds().H()/2))/2 - 40
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

