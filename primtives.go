package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

var (
	White = []float32{1, 1, 1, 1}
	Black = []float32{0, 0, 0, 1}
	Red   = []float32{1, 0, 0, 1}
	Green = []float32{0, 1, 0, 1}
	Blue  = []float32{0, 0, 1, 1}
)

func RGB(r uint, g uint, b uint) []float32 {
	return []float32{
		Map(float32(r), 0, 255, 0, 1),
		Map(float32(g), 0, 255, 0, 1),
		Map(float32(b), 0, 255, 0, 1),
		100,
	}
}

func RGBA(r uint, g uint, b uint, a uint) []float32 {
	return []float32{
		Map(float32(r), 0, 255, 0, 1),
		Map(float32(g), 0, 255, 0, 1),
		Map(float32(b), 0, 255, 0, 1),
		Map(float32(a), 0, 100, 0, 1),
	}
}

// TODO
func Hex(code string) {

}

// TODO refactor
func drawTriangle(x int, y int, width int, height int) {
	_x := Map(float32(x), 0, 1000, -1, 1)
	_y := Map(float32(y), 0, 1000, -1, 1)
	_width := Map(float32(height), 0, 1000, -1, 1)
	_height := Map(float32(width), 0, 1000, -1, 1)

	if _x >= 0 {
		_width = Map(float32(height), 0, 1000/2, 0, 1)
	}

	if _y > 0 {
		_width = Map(float32(height), 0, 1000/2, 0, 1)
	}

	vertices := []float32{
		_x, _height, 0, //Top left
		_width, _height, 0, //Top right
		_width, _y, 0, //Bottom right
		//-0.5, -0.5, 0, //Bottom left
	}

	vao := makeVAO(vertices)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/3))
}

func drawRect(x int, y int, width int, height int, colour []float32) {
	w, h := window.GetSize()
	xStartPos := Map(float32(x), 0, float32(w), -1, 1)
	xEndPos := Map(float32(x+width), 0, float32(w), -1, 1)
	yStartPos := Map(float32(y), 0, float32(h), 1, -1)
	yEndPos := Map(float32(y+height), 0, float32(h), 1, -1)

	vertices := []float32{
		//Left triangle
		xStartPos, yEndPos, colour[0], colour[1], colour[2], colour[3], //Top left
		xStartPos, yStartPos, colour[0], colour[1], colour[2], colour[3], //Bottom left
		xEndPos, yStartPos, colour[0], colour[1], colour[2], colour[3], //Bottom right

		//Right triangle
		xStartPos, yEndPos, colour[0], colour[1], colour[2], colour[3], //Top left
		xEndPos, yEndPos, colour[0], colour[1], colour[2], colour[3], //Top right
		xEndPos, yStartPos, colour[0], colour[1], colour[2], colour[3], //Bottom right
	}

	vao := makeVAO(vertices)
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/3))
}
