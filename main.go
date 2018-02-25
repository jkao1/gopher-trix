package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	screen := NewScreen()
	WriteScreenToPPM(screen)
}

func WriteScreenToPPM(screen [][][]int) {
	return
	file, err := os.OpenFile("foo.ppm", os.O_CREATE | os.O_WRONLY, 0644)

	if (err != nil) {
		panic(err)
	}

	defer file.Close()

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("P3 %d %d 255\n", XRES, YRES))
	// numerous string joins

	// Write buffer to file
	file.WriteString(buffer.String())
}
