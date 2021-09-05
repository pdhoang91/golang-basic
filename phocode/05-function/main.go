package main

import "fmt"

func main() {
	xs := []float64{98, 93, 82, 83}
	fmt.Println(average(xs))
	x := 5
	f(x)
	fmt.Println(f1())
	fmt.Println(f3())
	z, t := f4()
	fmt.Println(z, t)
	fmt.Println(add(1, 2, 3))
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))
	h := 0
	increment := func() int {
		h++
		return h
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(factorial(3))
	defer second()
	first()
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("Co loi xay ra")
}

func average(input []float64) float64 {
	total := 0.0
	for _, value := range input {
		total += value
	}
	return total / float64(len(input))
}

func f(x int) {
	fmt.Println(x)
}

func f1() int {
	return f2()
}

func f2() int {
	return 1
}

func f3() (r int) {
	r = 1
	return
}

func f4() (int, int) {
	return 5, 6
}

func add(argument ...int) int {
	total := 0
	for _, value := range argument {
		total += value
	}
	return total
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}
