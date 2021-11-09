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
		fmt.Println(a + b) // receive value from channel
	} else {
		fmt.Println("tried to read from closed channel")
	}
}

func main() {
	fmt.Printf("Start Main func at: %s\n", time.Now().Format(time.RFC1123))
	// Creating a channel
	ch := make(chan int)

	go sum(5, ch) // start a goroutine

	ch <- 25 // send value to channel

	close(ch)

	go sum(5, ch) // we try to receive data from a closed channel here, it does not block

	// ch <- 30 // this however results in panic!
	fmt.Printf("End Main func at: %s\n", time.Now().Format(time.RFC1123))

	// you can't see the message about closed channel unless you use something to wait until it finishes
	// uncomment to check it out
	//time.Sleep(time.Second * 10)
}
