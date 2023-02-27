package main

import (
	"fmt"
)

func main() {

	var arr [100]bool

	for i := 0; i < 100; i++ {
		arr[i] = false
	}
	fmt.Println(arr, "\n")
	var index = 1
	for index <= 100 {
		for i := index; i < 100; i += index {
			arr[i] = !arr[i]
		}
		index++
	}
	for i, state := range arr {
		if state {
			fmt.Println(i)
		}
	}
}
