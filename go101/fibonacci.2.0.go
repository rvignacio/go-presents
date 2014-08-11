// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func Fibonacci(msg string) {
	x, y := 0, 1
	for {
		x, y = y, x+y
		fmt.Printf("%s: %d!\n", msg, x)
		time.Sleep(time.Duration(rand.Intn(100)*10) * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	go Fibonacci("Fib")
}

// STOP OMIT
