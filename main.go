package main

import (
	"os"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/golang/freetype/truetype"
)

var width, height = 1000, 1000
var window *glfw.Window

func main() {
	runtime.LockOSThread()

	font, loadErr := os.ReadFile("fonts/Inter/static/Inter-Medium.ttf")
	truetype.Parse(font)

	err := glfw.Init()

	if err != nil {
		panic(err)
	}

	defer glfw.Terminate()

	window = InitGlfw(width, height, "Window")
	program := InitOpenGL()

	window.SetFramebufferSizeCallback(updateViewport)

	for !window.ShouldClose() {
		draw(window, program)
	}

}

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(1, 1, 1, 1)
	gl.UseProgram(program)

	sidebar := HStack{
		x: 0, y: 0,
		width: 500, height: 500,
		spacing: 20,
		color:   Black,
		children: []Widget{
			&VStack{
				x: 0, y: 0,
				width: 150, height: 150,
				spacing: 20,
				color:   Red,
			},
			&HStack{
				x: 0, y: 0,
				width: 150, height: 150,
				spacing: 20,
				color:   Green,
			},
		},
	}

	sidebar.render()

	glfw.PollEvents()
	window.SwapBuffers()
}

func updateViewport(window *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}
