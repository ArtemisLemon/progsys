package main

import "sync"
import "fmt"

var wg sync.WaitGroup
var mu sync.Mutex //RWMutex if you want to allow several reads at the same time

func writeMap(m map[int]int, k int) {
	defer wg.Done()
	mu.Lock()
	m[k] = k
	mu.Unlock()
}

func main() {

	m := make(map[int]int)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go writeMap(m, i)
	}

	wg.Wait()
	fmt.Println(m)
}
