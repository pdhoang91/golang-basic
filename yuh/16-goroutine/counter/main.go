package main

import (
	"fmt"
	"sync"
)

var wq = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {

	for i := 0; i < 10; i++ {
		wq.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}

	wq.Wait()
	fmt.Println("Done")

}

func sayHello() {
	fmt.Println("Hello ", counter)
	m.RUnlock()
	wq.Done()
}

func increment() {
	counter++
	m.Unlock()
	wq.Done()
}
