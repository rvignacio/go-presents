// +build OMIT

package main

import (
	"fmt"
)

// START OMIT
func adder() func(int) int {
	sum := 0 // HL
	return func(x int) int {
		sum += x // HL
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

//STOP OMIT
