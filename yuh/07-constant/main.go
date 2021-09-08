package main

import (
	"fmt"
)

const (
	red    = 0
	yellow = 1
	blue   = 2
)

const (
	x = iota
	y
	z
)

const (
	_ = 1<<iota*10 + 5
	a
	b
	c
)

func main() {
	const i int16 = 100
	const red = 10000
	fmt.Println("Hello I'm Tom")
	fmt.Printf("%v, %T\n", i, i)
	fmt.Println("Hello I'm Tom")
	fmt.Printf("%v, %T\n", red, red)

	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", z, z)

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
}
