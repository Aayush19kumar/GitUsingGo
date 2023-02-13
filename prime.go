package main

import "fmt"

func main() {
	fmt.Println("Enter number to check whether it is prime or not")
	var n int
	var count int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	if count == 2 {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not Prime")
	}
}
