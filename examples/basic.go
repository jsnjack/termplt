package main

import (
	"fmt"

	"github.com/jsnjack/termplt"
)

func main() {
	fmt.Println("Line Chart")
	chart := termplt.NewLineChart()
	chart.AddLine(
		[]float64{0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60},
		[]float64{0, 0.5, 0, 1.2, 0, 0, 5, 4, 1, 1.1, 0, 0, 0, 0},
		termplt.ColorBlue,
	)
	fmt.Println(chart.String())

	fmt.Println("Line Chart with axis")
	chart = termplt.NewLineChart()
	chart.AddLine(
		[]float64{0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60},
		[]float64{0, 0.5, 0, 1.2, 0, 0, 5, 4, 1, 1.1, 0, 0, 0, 0},
		termplt.ColorBlue,
	)
	chart.SetXLabel("mins")
	chart.SetYLabel("rain, mm")
	fmt.Println(chart.String())

	fmt.Println("Line Chart with time axis")
	chart = termplt.NewLineChart()
	chart.AddLine(
		[]float64{1633072800, 1633076400, 1633080000, 1633083600, 1633087200, 1633090800, 1633094400, 1633098000, 1633101600, 1633105200, 1633108800, 1633112400, 1633116000},
		[]float64{1, 2, 2, 1.2, 0, 0, 2, 3, 4, 5, 1, 0, 0, 1},
		termplt.ColorBlue,
	)
	chart.SetXLabelAsTime("time", "15:04")
	chart.SetYLabel("rain, mm")
	fmt.Println(chart.String())

	fmt.Println("Line Chart, no duplicated Y values")
	chart = termplt.NewLineChart()
	chart.AddLine(
		[]float64{0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60},
		[]float64{0, 0.1, 0.1, 0.2, 0.3, 0, 0, 0, 0, 0, 0.1, 0, 0, 0},
		termplt.ColorBlue,
	)
	chart.SetYLabel("rain, mm")
	fmt.Println(chart.String())
}
