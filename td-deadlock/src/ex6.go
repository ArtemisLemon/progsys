package main

import "fmt"

var x =10 //global choice of number of eXchanges

func ping(ready, next, done chan bool) {
	// for i := 0; i < 5; i++ { //can only send the ball back 5 times
	for i := 0; i < x; i++ {
		<-ready
		fmt.Println("Ping")
		next <- true
	}
	done <- true
}

func pong(ready, next, done chan bool) {
	// for i := 0; i < 10; i++ { //wants to send the ball back 10 times
	for i := 0; i < x; i++ {
		<-ready
		fmt.Println("Pong")
		next <- true
	}
	done <- true
}

func main() {
	pingchan := make(chan bool, 1)
	pongchan := make(chan bool, 1)
	done := make(chan bool, 2)

	pingchan <- true

	go ping(pingchan, pongchan, done)
	go pong(pongchan, pingchan, done)

	<-done
	<-done
}
