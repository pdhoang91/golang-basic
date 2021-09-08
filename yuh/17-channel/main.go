package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg = sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)

	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 22
		wg.Done()
	}()

	go func() {
		ch <- 42
		fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()
}
