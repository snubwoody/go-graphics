package main

import (
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const width, height = 1000, 1000

func main() {
	runtime.LockOSThread()

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

	rect := NewRow(0, 0, 250, 100, 10, RGB(0, 255, 100))
	rect3 := NewRow(0, 0, 250, 100, 10, RGB(0, 255, 100))
	rect4 := NewRow(0, 0, 250, 100, 10, RGB(0, 255, 100))
	rect2 := NewRow(0, 0, 500, 100, 10, RGB(0, 255, 100), rect3, rect4)

	box := NewColumn(0, 0, 250, 1000, 10, RGB(255, 255, 255), rect, rect2)
	box.render()

	glfw.PollEvents()
	window.SwapBuffers()
}
