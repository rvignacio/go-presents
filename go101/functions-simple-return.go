// +build OMIT

package main

import (
	"fmt"
)

// START OMIT
func Sum(x, y int) int {
	return x + y
}

func main() {
	var s int
	s = Sum(2, 4)
	fmt.Println(s)
}

// STOP OMIT
