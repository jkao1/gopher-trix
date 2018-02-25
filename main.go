package main

import (
	"fmt"
	"os"
)

var screen [][][]int

func main() {
	screen = NewScreen()
	fmt.Println(screen)
}

func WriteScreenToPPM() {
	file, err := os.OpenFile("foo.ppm", os.O_CREATE | os.O_WRONLY, 0644)

	if (err != nil) {
		panic(err)
	}

	var buffer bytes.Buffer

	buffer.WriteString("P3 %d %d 255\n", XRES, YRES)
	// numerous string joins

	// Write buffer to file
	file.WriteString(buffer.String())
}
