package main

import (
	"fmt"
	"strconv"
	"time"
)

// this channel only receives values, trying to send something results in compilation error
func printFromChannel(ch <-chan string) {
	fmt.Println(<-ch) // receive value from channel
	// ch <- "something" // unidirectional-example\main.go:12:5: invalid operation: ch <- "something" (send to receive-only type <-chan string)
}

// this channel only sends values, trying to receive something from it results in compilation error
func convert(a int, ch chan<- string) {
	ch <- strconv.Itoa(a) // send value to channel
	// fmt.Print(<-ch) // unidirectional-example\main.go:18:12: invalid operation: <-ch (receive from send-only type chan<- string)
}

func main() {
	// Creating a bidirectional channel
	ch := make(chan string)

	go printFromChannel(ch)
	go convert(10, ch) // start a goroutine

	time.Sleep(time.Second)
}
