package main

import "fmt"
import "sync"

var w sync.WaitGroup

func ping(done chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Println("Ping")
	}
	w.Done()
}

func main() {
	// done := make(chan bool, 1)
	w.Add(1)
	go ping(nil)
	// <-done //nothing in channel to remove, but waiting for this channel is how we wait for ping to print
}
