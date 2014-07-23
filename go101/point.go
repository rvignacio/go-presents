// +build OMIT

package main

import (
	"fmt"
)

// START OMIT
type Point struct {
	x, y int
}

func main() {
	p := Point{
		x: 1,
		y: 4,
	}

	fmt.Println(p)
}

// STOP OMIT
