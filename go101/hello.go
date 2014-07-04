package main

import (
	"fmt"
	"time"
)

// Main function that will be called to start our program.
func main() {
	say("Hello Gophers!")
}

// A simple function.
func say(msg string) {
	// A for loop
	for i := 0; i < 10; i++ {
		// Call an exported function from package fmt.
		fmt.Printf("%s %d\n", msg, i)
		time.Sleep(time.Second * 1)
	}
}
