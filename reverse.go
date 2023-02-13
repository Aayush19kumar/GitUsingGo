package main

import "fmt"

func main() {
	fmt.Println("Enter a no.")
	var n int
	fmt.Scan(&n)
	var ans = 0
	for n != 0 {
		r := n % 10
		ans = ans*10 + r
		n = n / 10
	}
	fmt.Println("The reverse no. is= ", ans)
}
