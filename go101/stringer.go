// +build OMIT

package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d : %d)", p.x, p.y)
}

// START OMIT
// Define the Stringer interface.
type Stringer interface {
	String() string
}

func Print(items ...Stringer) {
	// Iterate the items slice and call Stringer.String() method.
	for _, item := range items {
		fmt.Println(item.String())
	}
}

func main() {
	p1 := Point{x: 1, y: 4}
	// Create a new Point and use it right away.
	p2 := new(Point)

	Print(p1, p2)
}

// STOP OMIT
