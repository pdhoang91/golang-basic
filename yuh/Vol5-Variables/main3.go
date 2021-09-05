package main3

import (
	"fmt"
)

func main() {
	var i int
	i = 42
	var j int = 54
	k := 11

	fmt.Printf("%v, %T", i, i)
	fmt.Printf("%v, %T", j, j)
	fmt.Printf("%v, %T", k, k)
}
