package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	screen := NewScreen()

	DrawLine(screen, 0, 0, XRES - 1, YRES - 1)

	WriteScreenToPPM(screen)
}

func WriteScreenToPPM(screen [][][]int) {
	file, err := os.OpenFile("foo.ppm", os.O_CREATE | os.O_WRONLY, 0644)

	if (err != nil) {
		panic(err)
	}

	defer file.Close()

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("P3 %d %d 255\n", XRES, YRES))
	for i := 0; i < YRES; i++ {
		for j := 0; j < XRES; j++ {
			rgb := screen[i][j]
			buffer.WriteString(fmt.Sprintf("%d %d %d ", uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2])))
		}
	}

	file.WriteString(buffer.String())
}
