package main

type Widget interface {
	render()
	translate(x int, y int)
	size() (int, int)
}

type HStack struct {
	x, y          int
	width, height int
	spacing       int
	color         []float32
	children      []Widget
}

func (stack HStack) render() {
	drawRect(
		stack.x,
		stack.y,
		stack.width,
		stack.height,
		stack.color,
	)

	yOffset := stack.x

	for _, child := range stack.children {
		child.translate(0, yOffset)
		_, h := child.size()
		yOffset += h + stack.spacing
		child.render()
	}
}

func (stack HStack) translate(x int, y int) {

}

func (stack HStack) size() (int, int) {
	return stack.x, stack.y
}

type VStack struct {
	x, y          int
	width, height int
	spacing       int
	color         []float32
	children      []Widget
}

func (stack VStack) render() {
	drawRect(
		stack.x,
		stack.y,
		stack.width,
		stack.height,
		stack.color,
	)

	xOffset := stack.x

	for _, child := range stack.children {
		child.translate(xOffset, 0)
		w, _ := child.size()
		xOffset += w + stack.spacing
		child.render()
	}
}

func (stack VStack) translate(x int, y int) {

}

func (stack VStack) size() (int, int) {
	return stack.x, stack.y
}

type Frame struct {
	x, y                   int
	width, height          int
	xSpacing, ySpacing     int
	xDirection, yDirection bool
	color                  []float32
	children               []*Frame
}

func (frame Frame) render() {
	drawRect(
		frame.x,
		frame.y,
		frame.width,
		frame.height,
		frame.color,
	)

	xOffset := frame.x
	yOffset := frame.y

	for _, child := range frame.children {
		if frame.xDirection {
			child.translate(xOffset, 0)
		}

		if frame.yDirection {
			child.translate(0, yOffset)
		}
		xOffset += child.width + frame.xSpacing
		yOffset += child.height + frame.ySpacing
		child.render()
	}

}

func (frame *Frame) translate(x int, y int) {
	frame.x += x
	frame.y += y
}

func (frame Frame) size() (int, int) {
	return frame.x, frame.y
}
