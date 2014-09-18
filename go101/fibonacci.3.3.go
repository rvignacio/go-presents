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
	}()

	return c
}

// STARTFANIN OMIT
func fanIn(input1, input2 <-chan string, quit <-chan time.Time) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case val := <-input1:
				c <- val
			case val := <-input2:
				c <- val
			case <-quit:
				close(c)
				return
			}
		}
	}()
	return c
}

// STOPFANIN OMIT

// STARTMAIN OMIT
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

// STOPMAIN OMIT
