package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

// LineChart represents a line chart
type LineChart struct {
	canvas Canvas
	width  int
	height int
	lines  []line
}

// line represents a line in the LineChart
type line struct {
	x     []float64
	y     []float64
	color string
}

// Add a line to the chart
func (l *LineChart) AddLine(x []float64, y []float64, color string) {
	line := line{x: x, y: y, color: color}
	l.lines = append(l.lines, line)
}

// SetSize sets the size of the line chart
func (l *LineChart) SetSize(width, height int) {
	l.width = width
	l.height = height
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

func (l *LineChart) findMinY() float64 {
	min := l.lines[0].y[0]
	for _, line := range l.lines {
		for _, v := range line.y {
			if v < min {
				min = v
			}
		}
	}
	return min
}

func (l *LineChart) findMinX() float64 {
	min := l.lines[0].x[0]
	for _, line := range l.lines {
		for _, v := range line.x {
			if v < min {
				min = v
			}
		}
	}
	return min
}

func (l *LineChart) generateYLabels(numLines int, yPostfix string) []string {
	maxY := l.findMaxY()
	minY := l.findMinY()
	labels := make([]string, numLines)
	step := (maxY - minY) / float64(numLines)
	for i := 0; i < numLines; i++ {
		labels[i] = fmt.Sprintf("%.1f", minY+float64(i)*step)
	}
	if yPostfix != "" {
		labels[len(labels)-1] = yPostfix
	}
	maxLen := 0
	for _, label := range labels {
		labelR := []rune(label)
		if len(labelR) > maxLen {
			maxLen = len(labelR)
		}
	}
	for i := 0; i < len(labels); i++ {
		labels[i] = ensureLen(labels[i], maxLen)
	}
	// Reverse the order of the labels
	for i, j := 0, len(labels)-1; i < j; i, j = i+1, j-1 {
		labels[i], labels[j] = labels[j], labels[i]
	}
	return labels
}

func (l *LineChart) generateXLabels(xPostfix string) []string {
	maxX := l.findMaxX()
	minX := l.findMinX()
	labelCount := l.width/2 + 1
	axisAvailableChars := make([]string, labelCount)
	step := (maxX - minX) / float64(labelCount)
	stepValues := make([]float64, labelCount)
	for i := 0; i < labelCount; i++ {
		stepValues[i] = minX + float64(i)*step
	}
	minFreeSpace := 6
	for i := range axisAvailableChars {
		axisAvailableChars[i] = "⎺"
	}
	lastLabel := fmt.Sprintf("%.1f", maxX)
	for {
		if isXLabelsFull(axisAvailableChars, len(lastLabel)+minFreeSpace) {
			break
		}
		pos := 0
		for i := len(axisAvailableChars) - 1; i >= 0; i-- {
			if axisAvailableChars[i] != "⎺" {
				pos = i
				break
			}
		}
		if pos != 0 {
			pos += minFreeSpace
		}
		if pos >= len(axisAvailableChars) {
			break
		}
		axisAvailableChars = populateXLabelsSlice(axisAvailableChars, pos, fmt.Sprintf("%.1f", stepValues[pos]))
	}
	if xPostfix != "" {
		xPostfixLen := len([]rune(xPostfix))
		populateXLabelsSlice(axisAvailableChars, len(axisAvailableChars)-xPostfixLen, xPostfix)
		// make sure we have connecting ⎺ before the postfix
		for i := len(axisAvailableChars) - xPostfixLen - 1; i >= 0; i-- {
			if axisAvailableChars[i] == "⎺" {
				break
			} else {
				axisAvailableChars[i] = "⎺"
			}
		}
	}

	return axisAvailableChars
}

func populateXLabelsSlice(xLabels []string, pos int, label string) []string {
	runeLabel := []rune(label)
	for i, r := range runeLabel {
		xLabels[pos+i] = string(r)
	}
	return xLabels
}

func isXLabelsFull(xLabels []string, minLabelLen int) bool {
	// check if the last minLabelLen are ⎺
	for i := len(xLabels) - minLabelLen; i < len(xLabels); i++ {
		if xLabels[i] != "⎺" {
			return true
		}
	}
	return false
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

// StringWithAxis returns the string representation of the line chart with axis labels
func (l *LineChart) StringWithAxis(xPostfix, yPostfix string) string {
	data := l.String()
	splitted := strings.Split(data, "\n")
	splitted = splitted[:len(splitted)-1] // Remove the last empty line
	newData := ""
	yLabels := l.generateYLabels(len(splitted), yPostfix)
	for idx, line := range splitted {
		newData += fmt.Sprintf("%s⎹%s\n", yLabels[idx], line)
	}
	// Add the X axis
	paddingLen := len(yLabels[0])
	xLabels := l.generateXLabels(xPostfix)
	newData += strings.Repeat(" ", paddingLen+1) + strings.Join(xLabels, "") + "\n"
	return newData + "\n"
}

// Make a new line chart. The default size is auto-detected from the terminal size
func NewLineChart() LineChart {
	x, y, error := term.GetSize(int(os.Stdin.Fd()))
	if error != nil {
		x = 80
		y = 80
	}
	l := LineChart{
		width:  x,
		height: y,
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

// ensureLen ensures that the string has the specified length by padding it with spaces
func ensureLen(str string, length int) string {
	strLen := len([]rune(str))
	if strLen < length {
		return strings.Repeat(" ", length-strLen) + str
	}
	return str
}
