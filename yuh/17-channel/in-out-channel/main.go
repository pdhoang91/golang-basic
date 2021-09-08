package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg = sync.WaitGroup{}
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {
		// for i := range ch {
		// 	fmt.Println(i)
		// }
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}

		//ch <- 22
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 43
		ch <- 44
		close(ch)
		//fmt.Println("tom", i)
		wg.Done()
	}(ch)
	wg.Wait()
}
