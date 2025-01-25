package main

import (
	"fmt"
)

func main() {
	chart := NewLineChart()
	x := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	y := []float64{0, 0.5, 0, 1.2, 0, 0, 5, 4, 0, 0, 0, 0}
	yy := []float64{0, 0, 0, 1, 3, 14, 7, 6, 2, 0, 0, 0}
	// x := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// y := []float64{0, 1, 2, 2, 1, 0, 5, 4, 0, 0}
	chart.AddLine(x, yy, RedColor)
	chart.AddLine(x, y, BlueColor)
	// chart.SetSize(10, 4)
	fmt.Printf("%s", chart.StringWithAxis(YellowColor+"minutes"+ResetColor, CyanColor+"percip. mm"+ResetColor))
}
