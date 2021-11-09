package main

import (
	"fmt"
	"time"
)

func square(a int, ch chan int) {
	time.Sleep(time.Second * 5)

	ch <- a * a // send value to channel
}

func main() {
	fmt.Printf("Start Main func at: %s\n", time.Now().Format(time.RFC1123))
	// Creating a channel
	ch := make(chan int)

	go square(5, ch)                           // start a goroutine
	fmt.Printf("Result square is: %d\n", <-ch) // get value from channel
	fmt.Printf("End Main func at: %s\n", time.Now().Format(time.RFC1123))
}
