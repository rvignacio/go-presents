// +build OMIT

package main

import (
	"fmt"
	"time"
)

func Fibonacci(msg string) {
	x := 0
	y := 1
	for {
		x, y = y, x+y
		fmt.Printf("%s: %d!\n", msg, x)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	Fibonacci("Fib")
}
