package main

import "github.com/go-gl/gl/v4.1-core/gl"

type Frame struct {
	x        int
	y        int
	width    int
	height   int
	xSpacing int
	ySpacing int
	color    Color
	children []*Frame
}

func (frame Frame) render() {
	drawRect(
		frame.x,
		frame.y,
		frame.width,
		frame.height,
		frame.color.normalize(),
	)

	xOffset := frame.x
	yOffset := frame.y

	for _, child := range frame.children {
		child.translate(xOffset, yOffset)
		xOffset += child.width + frame.xSpacing
		yOffset += child.height + frame.ySpacing
		child.render()
	}

}

func (frame *Frame) position() (int, int) {
	return 0, 0
}

func (frame *Frame) translate(x int, y int) {
	frame.x += x
	frame.y += y
}

type Color [3]uint

func RGB(r uint, g uint, b uint) Color {
	return [3]uint{r, g, b}
}

func (color Color) normalize() [3]float32 {
	return [3]float32{
		Map(float32(color[0]), 0, 255, 0, 1),
		Map(float32(color[1]), 0, 255, 0, 1),
		Map(float32(color[2]), 0, 255, 0, 1),
	}
}

//TODO refactor
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

func drawRect(x int, y int, width int, height int, colour [3]float32) {
	xStartPos := Map(float32(x), 0, 1000, -1, 1)
	xEndPos := Map(float32(x+width), 0, 1000, -1, 1)
	yStartPos := Map(float32(y), 0, 1000, 1, -1)
	yEndPos := Map(float32(y+height), 0, 1000, 1, -1)

	vertices := []float32{
		//Left triangle
		xStartPos, yEndPos, colour[0], colour[1], colour[2], //Top left
		xStartPos, yStartPos, colour[0], colour[1], colour[2], //Bottom left
		xEndPos, yStartPos, colour[0], colour[1], colour[2], //Bottom right

		//Right triangle
		xStartPos, yEndPos, colour[0], colour[1], colour[2], //Top left
		xEndPos, yEndPos, colour[0], colour[1], colour[2], //Top right
		xEndPos, yStartPos, colour[0], colour[1], colour[2], //Bottom right
	}

	vao := makeVAO(vertices)
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/3))
}
