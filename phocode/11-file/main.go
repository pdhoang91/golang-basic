package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("D:/Study/Go/readme.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	str := string(bs)
	fmt.Println(str)
}
