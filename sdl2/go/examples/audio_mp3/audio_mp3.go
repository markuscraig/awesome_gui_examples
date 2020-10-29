package main

import (
	"flag"
	"log"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	mp3Path := flag.String("mp3", "./test.mp3", "mp3 file path")
	flag.Parse()

	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		log.Println(err)
		return
	}
	defer sdl.Quit()

	if err := mix.Init(mix.INIT_MP3); err != nil {
		log.Println(err)
		return
	}
	defer mix.Quit()

	if err := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Println(err)
		return
	}
	defer mix.CloseAudio()

	if music, err := mix.LoadMUS(*mp3Path); err != nil {
		log.Println(err)
	} else if err = music.Play(1); err != nil {
		log.Println(err)
	} else {
		sdl.Delay(5000)
		music.Free()
	}
}
