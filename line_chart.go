package main

type LineChart struct {
	canvas Canvas
	width  int
	height int
	lines  []Line
}

type Line struct {
	x     []float64
	y     []float64
	color string
}

// Add a line to the chart
func (l *LineChart) AddLine(x []float64, y []float64, color string) {
	line := Line{x: x, y: y, color: color}
	l.lines = append(l.lines, line)
}

func (l *LineChart) findMaxX() float64 {
	max := l.lines[0].x[0]
	for _, line := range l.lines {
		for _, v := range line.x {
			if v > max {
				max = v
			}
		}
	}
	return max
}

func (l *LineChart) findMaxY() float64 {
	max := l.lines[0].y[0]
	for _, line := range l.lines {
		for _, v := range line.y {
			if v > max {
				max = v
			}
		}
	}
	return max
}

// String returns the string representation of the line chart
func (l LineChart) String() string {
	l.canvas.Clear()

	// Resample data to fit the canvas
	for i, line := range l.lines {
		l.lines[i].x = resample(line.x, l.width)
		l.lines[i].y = resample(line.y, l.width)
	}

	// Find the maximum values for X and Y to scale the chart
	maxX := l.findMaxX()
	maxY := l.findMaxY()

	// Draw the lines
	for _, line := range l.lines {
		for i := 0; i < len(line.x); i++ {
			// Normalize the coordinates to fit the canvas
			nXi := int(line.x[i] * float64(l.width) / maxX)
			nYi := int(line.y[i] * float64(l.height) / maxY)
			// Invert Y axis to match the mathematical convention
			l.canvas.Set(nXi, l.height-nYi-1, line.color) // Invert Y axis
		}
	}
	return l.canvas.String()
}

// Make a new line chart
func NewLineChart() LineChart {
	l := LineChart{
		width:  80,
		height: 20,
	}
	l.canvas = NewCanvas()
	return l
}

// resample resizes the input slice to the new size using linear interpolation
func resample(input []float64, newSize int) []float64 {
	if len(input) == 0 || newSize == 0 {
		return []float64{}
	}

	output := make([]float64, newSize)
	scale := float64(len(input)-1) / float64(newSize-1)

	for i := range output {
		srcIndex := float64(i) * scale
		srcIndexInt := int(srcIndex)
		srcIndexFrac := srcIndex - float64(srcIndexInt)

		if srcIndexInt+1 < len(input) {
			output[i] = input[srcIndexInt]*(1-srcIndexFrac) + input[srcIndexInt+1]*srcIndexFrac
		} else {
			output[i] = input[srcIndexInt]
		}
	}

	return output
}
