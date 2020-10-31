# Go SDL2 Examples

Examples apps derived from the `go-sdl2-examples` repo:
* https://github.com/veandco/go-sdl2-examples/tree/master/examples

Example assets copied from the `go-sdl2-examples` repo:
* https://github.com/veandco/go-sdl2-examples/tree/master/assets

# Dependencies

### OSX
```
# install sdl2 libraries using homebrew
$ brew install sdl2{,_gfx,_image,_mixer,_net,_ttf}
```

**NOTE: Since homebrew no longer accepts options, use the following steps to reinstall `sdl2_mixer` with MP3 support:**

```
$ brew install mpg123
$ brew edit sdl2_mixer

# Change '--disable-music-mp3-mpg123' to '--enable-music-mp3-mpg123'
# Change '--disable-music-mp3-mpg123-shared' to '--enable-music-mp3-mpg123-shared'

$ brew reinstall --build-from-source sdl2_mixer 
```

# Go Modules

Go modules are used to include the `github.com/veandco/go-sdl2` dependencies.

# Running Examples using CLI

```
$ cd awesome_gui_examples/sdl2/go
$ go run examples/drawing-text/drawing-text.go -font ./assets/test.ttf -size 32
```

# Debugging Examples using VS-Code

Using VS-Code with the Golang extension allows you to debug each example using breakpoints, single-step, variable inspection, etc.

A `.vscode/launch.json` is included in the `go/sdl2` directory. Just open the `go/sdl2` directory in VS-Code:

```
$ code -n go/sdl2
```

Debug using VS-Code and the provided launch configurations:

![screenshot!](/common/screenshots/vscode_debug.png)
