package main

import (
	"fmt"
	"time"
)

func sum(a int, ch chan int) {
	time.Sleep(time.Second * 2)

	// you can check if you can receive data from channel using "ok" flag
	b, ok := <-ch
	if ok {
		fmt.Printf("a+b=%d\n", a+b) // receive value from channel
	} else {
		fmt.Println("tried to read from closed channel")
	}
}

func main() {
	fmt.Printf("Start Main func at: %s\n", time.Now().Format(time.RFC1123))
	// Creating a channel
	ch := make(chan int, 2)

	go sum(5, ch) // start a goroutine

	ch <- 25 // send value to channel
	ch <- 15 // send value to channel

	close(ch)

	// we try to receive data from a closed *buffered* channel here, it does not block
	// and since the buffer wasn't empty, we can read the remaining data from it!
	go sum(5, ch)

	// ch <- 30 // this however still results in panic!

	fmt.Printf("End Main func at: %s\n", time.Now().Format(time.RFC1123))

	// we wait until we finish reading from buffer
	time.Sleep(time.Second * 5)
}
