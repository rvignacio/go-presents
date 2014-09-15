// +build OMIT

package main

import (
	"fmt"
	"io"
)

func WriteString(w io.Writer, s string) (n int, err error) {
	/*...*/
	return
}

// START OMIT
func Split(sum int) (x, y int) { // HL
	x = sum * 4 / 9
	y = sum - x
	return // HL
}

func main() {
	a, b := Split(17) // HL
	fmt.Println(a, b)
}

// STOP OMIT
