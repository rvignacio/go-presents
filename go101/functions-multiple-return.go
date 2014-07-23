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
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := Split(17)
	fmt.Println(a, b)
}

// STOP OMIT
