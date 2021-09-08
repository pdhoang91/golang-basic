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
	}()

	go func() {
		ch <- 42
	}()
}
