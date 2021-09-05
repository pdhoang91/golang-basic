package main

import (
	"fmt"
)

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
	temp := new(int)
	one(temp)
	fmt.Println(&temp)
}

func zero(tmp *int) {
	*tmp = 0
}

func one(tmp *int) {
	*tmp = 1
}
