// Contains display functions for screens

package main

const XRES = 500
const YRES = 500

var DEFAULT_DRAW_COLOR []int = []int{255, 0, 0}

func NewScreen() (screen [][][]int) {
	screen = make([][][]int, YRES)

	for i, _ := range screen {
		screen[i] = make([][]int, XRES)

		for j, _ := range screen[i] {
			screen[i][j] = make([]int, 3)
		}
	}

	return
}

func DrawLine(screen [][][]int, x0, y0, x1, y1 int) {
	if x1 < x0 {
		x0, x1 = x1, x0
	}

	A := y1 - y0
	B := x0 - x1
	x := x0
	y := y0

	if B == 0 { // vertical line
		if y1 < y0 {
			y0, y1 = y1, y0
		}

		y = y0
		for y <= y1 {
			plot(screen, x, y)
			y++
		}

		return
	}

	slope := A / (-1.0 * B)
	var d int

	if slope >= 0 && slope <= 1 { // octant 1
		d = 2*A + B
		for x <= x1 && y <= y1 {
			plot(screen, x, y)
			if d > 0 {
				y++
				d += 2*B
			}
			x++
			d += 2*A
		}
	}

	if slope > 1 { // octant 2
		d = A + 2*B
		for x <= x1 && y <= y1 {
			plot(screen, x, y)
			if d < 0 {
				x++
				d += 2*A
			}
			y++
			d += 2*B
		}
	}

	if slope < 0 && slope >= -1 { // octant 8
		d = 2*A - B
    for x <= x1 && y >= y1 {
			plot(screen, x, y)
			if d < 0 {
				y--
				d -= 2*B
			}
			x++
			d += 2*A
		}
	}

	if slope < -1 { // octant 7
		d = A - 2*B
		for x <= x1 && y >= y1 {
			plot(screen, x, y)
			if d > 0 {
				x++
				d += 2*A
			}
			y--
			d -= 2*B
		}
	}
}

func plot(screen [][][]int, x, y int) {
	if x >= 0 && x < XRES && y >= 0 && y < YRES {
		screen[y][x] = DEFAULT_DRAW_COLOR[:]
	}
}
