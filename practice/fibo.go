package main

import (
	"fmt"
)

func main() {

	fmt.Print("enter number\n ")
	var n int
	fmt.Scan(&n)
	var a, b = 0, 1
	if n >= 2 {
		for i := 2; i <= n; i++ {
			temp := b
			b = a + b
			a = temp
		}
	}
	fmt.Println(b)
}
