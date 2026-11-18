package main

import (
	"log"
)


func prod(n int, com chan int) {
	defer close(com) //after this routine finishes com won't be written to
	for i := 0; i < n; i++ {
		com <- i
	}
}

func trans(com1 chan int, com2 chan int) {
	defer close(com2) //after this routine finishes com2 won't be written to
	for tmp := range com1{ //loop ends when com1 is closed
		com2<-tmp
	}
}

func cons(com chan int) {
	for tmp := range com{
		log.Print(tmp)
	}
}

func main() {
	var com1 chan int = make(chan int,5)
	var com2 chan int = make(chan int,5)

	go prod(1000, com1)
	go trans(com1,com2)
	cons(com2)
}
