package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg = sync.WaitGroup{}
	ch1 := make(chan int, 50)
	ch2 := make(chan int, 50)
	wg.Add(2)

	go func() {

		for {
			select {
			case i := <-ch1:
				fmt.Printf("Channel 1: %v\n", i)
			case j := <-ch2:
				fmt.Printf("Channel 1: %v\n", j)
			default:
				break
			}
		}

		//ch <- 22
		wg.Done()
	}()

	go func() {
		ch1 <- 42
		ch1 <- 43
		close(ch1)
		ch2 <- 44
		close(ch2)
		//fmt.Println("tom", i)
		wg.Done()
	}()
	wg.Wait()
}
