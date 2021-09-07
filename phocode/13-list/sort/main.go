package main

import (
	"fmt"
	"sort"
)

func main() {

	x := []int{93, 48, 27, 784, 13}
	fmt.Println(x)
	sort.IntSlice.Sort(x)
	fmt.Println(x)

	y := []float64{3.5, 7.6, 8.93, 5.23, 7.609}
	fmt.Println(y)
	sort.Float64Slice.Sort(y)
	fmt.Println(y)
}
