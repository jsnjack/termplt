package main

import (
	"fmt"
)

func main() {
	chart := NewLineChart()
	x := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	y := []float64{0, 0.5, 0, 1.2, 0, 0, 5, 4, 0, 0}
	yy := []float64{0, 0, 0, 1, 3, 10, 7, 6, 2, 0}
	// x := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// y := []float64{0, 1, 2, 2, 1, 0, 5, 4, 0, 0}
	chart.AddLine(x, y, BlueColor)
	chart.AddLine(x, yy, RedColor)
	fmt.Printf("%s", chart.String())
}
