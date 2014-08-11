// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func Fibonacci(msg string) <-chan string {
	c := make(chan string)
	go func() {
		x, y := 0, 1
		for {
			x, y = y, x+y
			c <- fmt.Sprintf("%s: %d!", msg, x)
			time.Sleep(time.Duration(rand.Intn(100)*10) * time.Millisecond)
		}
	}()

	return c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	alice := Fibonacci("Alice")
	bob := Fibonacci("Bob  ")
	for i := 0; i < 5; i++ {
		fmt.Println(<-alice)
		fmt.Println(<-bob)
	}
	fmt.Println("I'm bored, good bye.")
}

// STOP OMIT
