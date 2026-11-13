package main

import "fmt"
var turn chan bool = make(chan bool,1)

func routine1() {
	for {
		if <-turn { //if this routine receives 
			fmt.Println("Routine 1") //do work
		}
		turn <-false //weather we worked, or <-false, we put the same message back in the channel
	}
}

func routine2() {
	for {
		if !<-turn { //if this routine receives 
			fmt.Println("Routine 2") //do work
		}
		turn <-true
	}
}

func main() {
	turn<-false
	go routine1()
	routine2()
}
