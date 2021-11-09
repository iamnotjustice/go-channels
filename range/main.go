package main

import (
	"fmt"
	"math/rand"
)

func sum(in chan int, out chan int) {
	sum := 0

	// we iterate over channel contents until it is closed
	for a := range in {
		sum += a
		fmt.Printf("Value from channel: %d\n", a)
	}

	out <- sum
}

// same thing, but without range
func sumNoRange(in, out chan int) {
	sum := 0

	for {
		a, ok := <-in
		if !ok {
			break
		}
		sum += a
		fmt.Printf("Value from channel: %d\n", a)
	}

	out <- sum
}

func generate(num int, ch chan int) {
	for i := 0; i < num; i++ {
		ch <- rand.Intn(10)
	}

	// closing a channel when finished writing to it
	close(ch)
}

func main() {
	// Creating a channel
	in := make(chan int)
	out := make(chan int)

	go generate(10, in)
	go sum(in, out)
	//go sumNoRange(in, out)

	// here we wait until we got our result from sum() goroutine
	fmt.Printf("Result = %d", <-out)
}
