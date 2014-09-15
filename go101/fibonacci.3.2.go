// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Fibonacci(msg string, howMany int) <-chan string {
	c := make(chan string)
	go func() {
		x, y := 0, 1
		for i := 0; i < howMany; i++ {
			x, y = y, x+y
			c <- fmt.Sprintf("%s: %d!", msg, x)
			time.Sleep(time.Duration(rand.Intn(100)*10) * time.Millisecond)
		}
		close(c)
	}()

	return c
}

// START OMIT
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	var closed1, closed2 bool
	go func() {
		for {
			if closed1 && closed2 {
				close(c)
				return
			}
			select {
			case val, ok := <-input1:
				if !ok {
					closed1 = true
				} else {
					c <- val
				}
			case val, ok := <-input2:
				if !ok {
					closed2 = true
				} else {
					c <- val
				}
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := fanIn(Fibonacci("Alice", 5), Fibonacci("Bob  ", 5))
	for {
		msg, ok := <-c
		if !ok {
			break
		}
		fmt.Println(msg)
	}
}

// STOP OMIT
