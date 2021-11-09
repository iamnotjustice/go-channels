package main

import (
	"fmt"
	"time"
)

func sum(a int, ch chan int) {
	time.Sleep(time.Second * 5)

	fmt.Println(a + <-ch) // receive value from channel
}

func main() {
	fmt.Printf("Start Main func at: %s\n", time.Now().Format(time.RFC1123))
	// Creating a channel
	ch := make(chan int)

	go sum(5, ch) // start a goroutine

	ch <- 25 // send value to channel
	fmt.Printf("End Main func at: %s\n", time.Now().Format(time.RFC1123))
}
