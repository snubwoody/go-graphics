package main

import (
	"log"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func InitGlfw(width int, height int, title string) *glfw.Window {
	glfw.WindowHint(glfw.Resizable, glfw.False)
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

	frag, err := os.ReadFile("triangle.frag")
	if err != nil {
		panic(err)
	}

	vert, err := os.ReadFile("triangle.vert")
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

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)
	return program
}
