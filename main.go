package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)

	// it is important to pass the wait group by reference
	// otherwise, the counter of the `wg` instance wouldn't decrement
	// when the goroutine has invoked wg.Done()
	go Hello(&wg)
	go Bye(&wg)

	wg.Wait()
}

func Hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello, world!")
}

func Bye(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("bye!")
}
