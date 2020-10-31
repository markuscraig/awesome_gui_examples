# Go GLFW Examples

Go examples using the GLFW framwork and the Go `github.com/gl-gl/glfw` module.

For OpenGL-ES2 targets (most Raspberry Pi boards), the Go glfw module needs to be fetched with the `gles2` tag:

```
$ go get -u -tags=gles2 github.com/go-gl/glfw/v3.3/glfw
```

# Go Modules

Go modules are used to include the `github.com/go-gl/glfw/v3.3/glfw` dependencies.

# Running Examples using CLI

```
$ cd awesome_gui_examples/glfw/go
$ go run examples/hello-world/hello-world.go -width 800 -height 600
```

# Debugging Examples using VS-Code

Using VS-Code with the Golang extension allows you to debug each example using breakpoints, single-step, variable inspection, etc.

A `.vscode/launch.json` is included in the `go/glfw` directory. Just open the `go/glfw` directory in VS-Code:

```
$ code -n go/glfw
```

Debug using VS-Code and the provided launch configurations:

![screenshot!](/common/screenshots/vscode_debug.png)
