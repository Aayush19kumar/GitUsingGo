package main

import "fmt"

func main() {
	fmt.Println("Enter no.")
	var n int
	fmt.Scan(&n)
	var ans int
	for n != 0 {
		r := n % 10
		ans += r
		n = n / 10
	}
	fmt.Println("The sum of the digits is :", ans)
}
