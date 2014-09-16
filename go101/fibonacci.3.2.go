// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Fibonacci(msg string) <-chan string {
	c := make(chan string)
	go func() {
		x, y := 0, 1
		for {
			x, y = y, x+y
			c <- fmt.Sprintf("%s: %d!", msg, x)
			time.Sleep(time.Duration(rand.Intn(100)*10) * time.Millisecond)
		}
		return
	}()

	return c
}

// START OMIT
func fanIn(input1, input2 <-chan string, quit <-chan time.Time) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case c <- <-input1:
			case c <- <-input2:
			case <-quit:
				close(c)
				return
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := fanIn(Fibonacci("Alice"), Fibonacci("Bob  "), time.After(5*time.Second))
	for {
		msg, ok := <-c
		if !ok {
			break
		}
		fmt.Println(msg)
	}
	fmt.Println("I'm bored, good bye.")
}

// STOP OMIT
