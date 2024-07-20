package main

import (
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var width, height = 1000, 1000
var window *glfw.Window

func main() {
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	defer glfw.Terminate()

	window = CreateWindow(width, height, "Window")
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
	drawText(500, 500, 50, 50, Black)

	glfw.PollEvents()
	window.SwapBuffers()
}

func updateViewport(window *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}
