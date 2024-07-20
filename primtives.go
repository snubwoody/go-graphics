package main

import (
	"image"
	"log"
	"os"

	_ "image/png"

	"github.com/fogleman/gg"
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

func drawText(x int, y int, width int, height int, colour []float32) {
	dc := gg.NewContext(1000, 1000)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	if fontErr := dc.LoadFontFace("fonts/Inter/static/Inter-Regular.ttf", 96); fontErr != nil {
		log.Fatal(fontErr)
	}
	dc.DrawStringAnchored("Hello world", 500, 500, 0.5, 0.5)
	dc.SavePNG("out.png")

	w, h := window.GetSize()
	xStartPos := Map(float32(x), 0, float32(w), -1, 1)
	xEndPos := Map(float32(x+width), 0, float32(w), -1, 1)
	yStartPos := Map(float32(y), 0, float32(h), 1, -1)
	yEndPos := Map(float32(y+height), 0, float32(h), 1, -1)
	_width := Map(float32(width), 0, float32(w), -1, 1)
	_height := Map(float32(height), 0, float32(h), -1, 1)

	reader, err := os.Open("out.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()

	image, _, err := image.Decode(reader)

	vertices := []float32{
		//Left triangle
		xStartPos, yEndPos, colour[0], colour[1], colour[2], colour[3], 0, 1, //Top left
		xStartPos, yStartPos, colour[0], colour[1], colour[2], colour[3], 0, 0, //Bottom left
		xEndPos, yStartPos, colour[0], colour[1], colour[2], colour[3], 1, 0, //Bottom right

		//Right triangle
		xStartPos, yEndPos, colour[0], colour[1], colour[2], colour[3], 0, 1, //Top left
		xEndPos, yEndPos, colour[0], colour[1], colour[2], colour[3], 1, 1, //Top right
		xEndPos, yStartPos, colour[0], colour[1], colour[2], colour[3], 1, 0, //Bottom right
	}

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(_width), int32(_height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(image))

	vao := makeVAO(vertices)
	gl.BindVertexArray(vao)
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.Ptr(0))
}
