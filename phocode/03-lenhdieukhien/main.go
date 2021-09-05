package main

import "fmt"

func main() {
	var i int = 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j <= 10; j++ {
		if i%2 == 0 {
			fmt.Println(j, "chan")
		} else {
			fmt.Println(j, "le")
		}
	}
}
