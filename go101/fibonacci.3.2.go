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

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { for { c <- <-input1 } }()
	go func() { for { c <- <-input2 } }()
	return c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := fanIn(Fibonacci("Alice"), Fibonacci("Bob  "))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("I'm bored, good bye.")
}

// STOP OMIT
