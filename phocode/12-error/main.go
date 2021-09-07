package main

import (
	"errors"
	"fmt"
)

func main() {
	arg1 := 10
	arg2 := 0
	result, err := divige(arg1, arg2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func divige(arg1, arg2 int) (int, error) {
	if arg2 == 0 {
		return -1, errors.New("Can't devige by 0")
	}
	return arg1 / arg2, nil
}
