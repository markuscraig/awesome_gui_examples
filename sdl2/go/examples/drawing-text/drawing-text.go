package main

import (
	"flag"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	fontPath string
	fontSize int
)

func run() (err error) {
	flag.StringVar(&fontPath, "font", "./test.ttf", "font file path")
	flag.IntVar(&fontSize, "size", 32, "font size")
	flag.Parse()

	var window *sdl.Window
	var font *ttf.Font
	var surface *sdl.Surface
	var text *sdl.Surface

	if err = ttf.Init(); err != nil {
		return
	}
	defer ttf.Quit()

	if err = sdl.Init(sdl.INIT_VIDEO); err != nil {
		return
	}
	defer sdl.Quit()

	// Create a window for us to draw the text on
	if window, err = sdl.CreateWindow("Drawing text", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN); err != nil {
		return
	}
	defer window.Destroy()

	if surface, err = window.GetSurface(); err != nil {
		return
	}

	// Load the font for our text
	if font, err = ttf.OpenFont(fontPath, fontSize); err != nil {
		return
	}
	defer font.Close()

	// Create a red text with the font
	if text, err = font.RenderUTF8Blended("Howdy World!", sdl.Color{R: 0, G: 255, B: 255, A: 255}); err != nil {
		return
	}
	defer text.Free()

	// Draw the text around the center of the window
	if err = text.Blit(nil, surface, &sdl.Rect{X: 400 - (text.W / 2), Y: 300 - (text.H / 2), W: 0, H: 0}); err != nil {
		return
	}

	// Update the window surface with what we have drawn
	window.UpdateSurface()

	// Run infinite loop until user closes the window
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		sdl.Delay(16)
	}

	return
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}
