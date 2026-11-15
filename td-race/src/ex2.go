package main

import (
	"log"
	"sync"
)

// var com1, com2 []int
var w sync.WaitGroup

func prod(n int, com chan int) {
	// defer w.Done()
	for i := 0; i < n; i++ {
		// com1 = append(com1, i)
		com <- i
	}
}

func trans(com1 chan int, com2 chan int) {
	// defer w.Done()
	for {
		// if len(com1) > 0 {
		// 	com2 = append(com2, com1[0])
		// 	com1 = com1[1:]
		// }
		tmp:=<-com1
		com2<-tmp
	}
}

func cons(com chan int) {
	// defer w.Done()
	for {
		// if len(com2) > 0 {
		// 	log.Print(com2[0])
		// 	com2 = com2[1:]
		// }
		log.Print(<-com)
	}
}

func main() {
	var com1 chan int = make(chan int,5)
	var com2 chan int = make(chan int,5)

	// w.Add(3)
	go prod(1000, com1)
	go trans(com1,com2)
	go cons(com2)
	// w.Wait()
}
