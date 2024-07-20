package main

import (
	"log"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func CreateWindow(width int, height int, title string) *glfw.Window {
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	return window
}

func InitOpenGL() uint32 {
	err := gl.Init()
	if err != nil {
		panic(err)
	}

	frag, err := os.ReadFile("shaders/triangle.frag")
	if err != nil {
		panic(err)
	}

	vert, err := os.ReadFile("shaders/triangle.vert")
	if err != nil {
		panic(err)
	}

	vertexShader, err := compileShader(strings.Join([]string{string(vert), "\x00"}, " "), gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(strings.Join([]string{string(frag), "\x00"}, " "), gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	gl.Viewport(0, 0, int32(width), int32(height))

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	gl.Enable(gl.BLEND)
	gl.Enable(gl.FRAMEBUFFER_SRGB)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)
	return program
}
