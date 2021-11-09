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

	fmt.Println("finished sum operation")
}

func generate(num int, ch chan int) {
	for i := 0; i < num; i++ {
		ch <- rand.Intn(10)
	}

	// closing a channel when finished writing to it
	close(ch)

	fmt.Println("finished generating numbers")
}

func main() {
	// Creating a buffered channel
	in := make(chan int, 5)
	out := make(chan int)

	// this should not fill up our buffer, so it won't block
	go generate(5, in)
	go sum(in, out)

	// here we wait until we got our result from sum() goroutine
	fmt.Printf("Result = %d", <-out)
}
