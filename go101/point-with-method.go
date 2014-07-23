// +build OMIT

package main

import (
	"fmt"
)

// START OMIT
type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d : %d)", p.x, p.y)
}

func main() {
	// p.x is implicit, it is initialized to its zero value
	p := Point{
		y: 4,
	}

	fmt.Println(p.String())
}

// STOP OMIT

// REFERENCE OMIT
// p is a Point reference so that the SetX() method can overwrite the caller.
func (p *Point) SetX(x int) {
	p.x = x
}

// STOPREFERENCE OMIT
