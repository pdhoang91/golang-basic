package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wq = sync.WaitGroup{}
	wq.Add(2)
	go func() {
		count("sheep")
		wq.Done()
	}()

	go func() {
		count("fish")
		wq.Done()
	}()
	wq.Wait()
	fmt.Println("Done")
	//time.Sleep(time.Second * 5)

}

func count(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
}
