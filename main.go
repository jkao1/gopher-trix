package main

func main() {
	screen := NewScreen()
	m := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 1, 1},
	}

	DrawLines(m, screen)
}

/*
func main() {
	screen := NewScreen()
	defer WriteScreenToPPM(screen)

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
*/
