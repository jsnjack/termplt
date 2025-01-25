package main

var pixel_map = [4][2]int{
	{0x1, 0x8},
	{0x2, 0x10},
	{0x4, 0x20},
	{0x40, 0x80}}

// Braille chars start at 0x2800
var braille_char_offset = 0x2800

func getPixel(y, x int) int {
	var cy, cx int
	if y >= 0 {
		cy = y % 4
	} else {
		cy = 3 + ((y + 1) % 4)
	}
	if x >= 0 {
		cx = x % 2
	} else {
		cx = 1 + ((x + 1) % 2)
	}
	return pixel_map[cy][cx]
}

type Canvas struct {
	chars  map[int]map[int]int
	colors map[int]map[int]string
}

// Make a new canvas
func NewCanvas() Canvas {
	c := Canvas{}
	c.Clear()
	return c
}

func (c Canvas) MaxY() int {
	max := 0
	for k := range c.chars {
		if k > max {
			max = k
		}
	}
	return max * 4
}

func (c Canvas) MinY() int {
	min := 0
	for k := range c.chars {
		if k < min {
			min = k
		}
	}
	return min * 4
}

func (c Canvas) MaxX() int {
	max := 0
	for _, v := range c.chars {
		for k := range v {
			if k > max {
				max = k
			}
		}
	}
	return max * 2
}

func (c Canvas) MinX() int {
	min := 0
	for _, v := range c.chars {
		for k := range v {
			if k < min {
				min = k
			}
		}
	}
	return min * 2
}

// Clear all pixels
func (c *Canvas) Clear() {
	c.chars = make(map[int]map[int]int)
	c.colors = make(map[int]map[int]string)
}

// Convert x,y to cols, rows
func (c Canvas) get_pos(x, y int) (int, int) {
	return (x / 2), (y / 4)
}

// Set a pixel of c with color
func (c *Canvas) Set(x, y int, color string) {
	px, py := c.get_pos(x, y)
	if m := c.chars[py]; m == nil {
		c.chars[py] = make(map[int]int)
	}
	if m := c.colors[py]; m == nil {
		c.colors[py] = make(map[int]string)
	}
	val := c.chars[py][px]
	mapv := getPixel(y, x)
	c.chars[py][px] = val | mapv
	c.colors[py][px] = color
}

// Retrieve the rows from a given view
func (c Canvas) Rows(minX, minY, maxX, maxY int) []string {
	minrow, maxrow := minY/4, (maxY)/4
	mincol, maxcol := minX/2, (maxX)/2

	ret := make([]string, 0)
	for rownum := minrow; rownum < (maxrow + 1); rownum = rownum + 1 {
		row := ""
		for x := mincol; x < (maxcol + 1); x = x + 1 {
			char := c.chars[rownum][x]
			color := c.colors[rownum][x]
			row += color + string(rune(char+braille_char_offset)) + ResetColor
		}
		ret = append(ret, row)
	}
	return ret
}

// Retrieve a string representation of the frame at the given parameters
func (c Canvas) Frame(minX, minY, maxX, maxY int) string {
	var ret string
	for _, row := range c.Rows(minX, minY, maxX, maxY) {
		ret += row
		ret += "\n"
	}
	return ret
}

func (c Canvas) String() string {
	return c.Frame(c.MinX(), c.MinY(), c.MaxX(), c.MaxY())
}
