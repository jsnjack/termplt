package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	width, height := 80, 80
	if len(os.Args) > 2 {
		w, err1 := strconv.Atoi(os.Args[1])
		h, err2 := strconv.Atoi(os.Args[2])
		if err1 == nil && err2 == nil {
			width, height = w, h
		}
	}

	canvas := NewCanvas()
	drawParabola(&canvas, width, height, "\033[31m") // Red color
	drawLinear(&canvas, width, height, "\033[34m")   // Blue color
	fmt.Println(canvas.String())
}

func drawParabola(c *Canvas, width, height int, color string) {
	scaleX := float64(width) / 2

	for x := -width / 2; x <= width/2; x++ {
		y := (x * x) / int(scaleX)
		c.Set(x+width/2, height/2-y, color)
	}
}

func drawLinear(c *Canvas, width, height int, color string) {
	for x := -width / 2; x <= width/2; x++ {
		y := x
		c.Set(x+width/2, height/2-y, color)
	}
}
