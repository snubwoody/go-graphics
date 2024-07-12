package main

import (
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const width, height = 1000, 1000

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	defer glfw.Terminate()

	window := InitGlfw(width, height, "Window")
	program := InitOpenGL()

	for !window.ShouldClose() {
		draw(window, program)
	}

}

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	box1 := Frame{0, 0, 100, 100, Color{255, 255, 255}}

	box := Row{
		x:        0,
		y:        0,
		width:    300,
		height:   1000,
		children: []Drawable{box1},
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
