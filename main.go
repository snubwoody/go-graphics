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

	rect := Frame{
		0, 0,
		100, 100,
		0, 0,
		Color{25, 205, 55},
		nil,
	}

	rect2 := Frame{
		0, 0,
		100, 100,
		0, 0,
		Color{5, 25, 25},
		nil,
	}

	box := Frame{
		50, 50,
		250, 1000,
		0, 0,
		RGB(255, 255, 255),
		[]*Frame{&rect, &rect2},
	}

	box.render()

	glfw.PollEvents()
	window.SwapBuffers()
}
