package main

import (
	"fmt"
	"math/rand"
	"time"
)

func routine(c chan int) {
	wait := rand.Intn(1000) 
	time.Sleep(time.Duration(wait) * time.Millisecond) //wait random amount of second
	c <- wait //send wait time
}

func main() {

	rand.Seed(int64(time.Now().Nanosecond())) //seed random

	c1 := make(chan int)
	c2 := make(chan int)

	go routine(c1)
	go routine(c2)

	for {
		select { //wait for a channel response
			case w := <-c1: //the case where 2 gets a value from channel 1 first
			fmt.Println("La première goroutine a été la plus rapide.")
			fmt.Println("Elle a mis", w, "millisecondes.")
			return
			case w := <-c2: //the case where 2 gets a value from channel 2 first
			fmt.Println("La deuxième goroutine a été la plus rapide.")
			fmt.Println("Elle a mis", w, "millisecondes.")
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timeout")
			return
		default:
			fmt.Println("skip select") //useful to not block a loop
		}
	}

}
