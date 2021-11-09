package main

import "fmt"

func f(c chan int) {
	fmt.Println(<-c)
}

func main() {
	fmt.Println("Started")
	var c chan int

	go f(c)

	c <- 5 // sending in nil-channel blocks forever

	select {} // empty select blocks forever as well!

	fmt.Print("Finished")
}
