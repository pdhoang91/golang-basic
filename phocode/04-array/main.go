package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total = total + x[i]
	}
	fmt.Println(total / 5)

	for _, value := range x {
		total = total + value
	}
	fmt.Println(total / float64(len(x)))

	y := [5]float64{98, 93, 77, 82, 83}
	fmt.Println(y)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	k := make(map[string]int)
	k["key"] = 10
	fmt.Println(k["key"])
}
