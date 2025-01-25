package main

import (
	"fmt"

	"github.com/jsnjack/termplt"
)

func main() {
	chart := termplt.NewLineChart()
	x := []float64{0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60}
	y := []float64{0, 0.5, 0, 1.2, 0, 0, 5, 4, 1, 1.1, 0, 0, 0, 0}
	chart.AddLine(x, y, termplt.ColorBlue)
	chart.SetYLabel("rain, mm")
	chart.SetXLabel("mins")
	fmt.Println(chart.String())
}
