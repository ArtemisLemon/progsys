package main

import "fmt"

func main() {

	var c chan int = make(chan int, 3)

	c <- 1 //send 1 to channel
	c <- 2
	c <- 3
	// c <- 4 // blocks channel

	fmt.Println(<-c) // receive message from channel
	// c <- 4 // channel slot has been freed to it works
	fmt.Println(<-c)
	fmt.Println(<-c)
	// fmt.Println(<-c) //reading empty channel is blocking
	// fmt.Println(<-c)

}
