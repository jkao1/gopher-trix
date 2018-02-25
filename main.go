package main

import (
	"fmt"
)

var screen [][][]int

func main() {
	screen = NewScreen()
	fmt.Println(screen)
}
