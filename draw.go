package main

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func NormalizeScreenCoordinates(x int, y int) (float32, float32) {
	monitor := glfw.GetPrimaryMonitor()
	if monitor == nil {
		log.Fatalln("Couldn't identify monitor")
	}

	videoMode := monitor.GetVideoMode()
	if videoMode == nil {
		log.Fatalln("Failed to get video mode")
	}

	normalizedXCoord := 2.0*((float32(x)-0.0)/(1000.0-0.0)) - 1.0
	normalizedYCoord := 2.0*((float32(y)-0.0)/(1000.0-0.0)) - 1.0

	return float32(normalizedXCoord), float32(normalizedYCoord)
}

// Map a value from one range to another
func Map(value float32, inMin float32, inMax float32, outMin float32, outMax float32) float32 {
	scale := (outMax - outMin) / (inMax - inMin)
	offset := inMin*((outMax-outMin)/(inMax-inMin)) + outMin
	output := value*scale + offset

	return output
}
