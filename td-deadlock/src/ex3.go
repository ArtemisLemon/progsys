package main

import (
	"fmt"
	"sync"
	"time"
)

var x, y int
var w sync.WaitGroup
var mx, my sync.Mutex

func switchxy() {
	mx.Lock()
	time.Sleep(time.Second) //forces deadlock when mutex lock order different
	my.Lock()
	for i := 0; i < 1000; i++ {
		x, y = y, x
	}
	mx.Unlock()
	my.Unlock()
	w.Done()
}

func switchyx() {
	// my.Lock()
	// time.Sleep(time.Second) //forces deadlock when mutex lock order different
	mx.Lock()
	my.Lock()
	for i := 0; i < 1000; i++ {
		x, y = y, x
	}
	my.Unlock()
	mx.Unlock()
	w.Done()
}

func main() {
	x = 5
	y = 7
	w.Add(2)
	go switchxy()
	go switchyx()

	//IF
		//switchXy gets mutex X
		//switchYx gets mutex Y
		//switchXy can't get mutex Y
		//switchYx can't get mutex X
	//Then deadlock
	
	//But it's possible for switchXy to get both mutex before switchYx asks for one (or vice versa)
	//Dodge deadlock

	//to avoid deadlock
	//ask for mutex in the same order for all routines

	w.Wait()
	fmt.Println("x vaut", x, "et y vaut", y)
}
