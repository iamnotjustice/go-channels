package main

import (
	"fmt"
	"time"
)

func main() {
	fastout := make(chan int)
	slowout := make(chan int)

	in := make(chan string)

	// these two anon-functions do the same action but one is a bit slower
	go func(num int, out chan<- int) {
		result := num * num
		time.Sleep(10 * time.Millisecond)
		out <- result
	}(10000, fastout)

	go func(num int, out chan<- int) {
		result := num * num
		time.Sleep(11 * time.Millisecond)
		out <- result
	}(2, slowout)

	// this one should be ready to receive before two computations above
	go func(in <-chan string) {
		fmt.Printf("%s\n", <-in)
	}(in)

	finished := false

	for !finished {
		// select takes whichever is ready to be received or sent first and does it's case
		// our "received" case should finish be ready first,
		// then one of the computations, fast or slow.
		// Play with time.Sleep numbers to check how this changes the order and how it affects the select-case.
		select {
		case res := <-fastout:
			fmt.Println("fast finished first, result:", res)
			finished = true
		case res := <-slowout:
			fmt.Println("slow finished first, result:", res)
			finished = true

		case in <- "received":
		// do nothing
		default:
			fmt.Printf("waiting\n")
		}

		time.Sleep(500 * time.Nanosecond)
	}
}
