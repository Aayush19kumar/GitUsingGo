package main

import "fmt"

func main() {

	var a = 5
	var b = 6
	fmt.Println("No. without swapping ", a, b)

	fmt.Print(swap(a, b))

}

func swap(a int, b int) (int, int) {

	a = a + b
	b = a - b
	a = a - b
	return a, b

}
