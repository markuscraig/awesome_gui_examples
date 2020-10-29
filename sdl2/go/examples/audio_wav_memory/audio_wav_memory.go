package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	wavFile := flag.String("wav", "./test.wav", "wav audio file")
	flag.Parse()

	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		log.Println(err)
		return
	}
	defer sdl.Quit()

	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Println(err)
		return
	}
	defer mix.CloseAudio()

	// Load entire WAV data from file
	data, err := ioutil.ReadFile(*wavFile)
	if err != nil {
		log.Println(err)
	}

	// Load WAV from data (memory)
	chunk, err := mix.QuickLoadWAV(data)
	if err != nil {
		log.Println(err)
	}
	defer chunk.Free()

	// Play 4 times
	chunk.Play(1, 3)

	// Wait until it finishes playing
	for mix.Playing(-1) == 1 {
		sdl.Delay(16)
	}
}
