package main

import "fmt"

func main() {
	fmt.Println("Enter to get factorial\n")

	var n int
	var ans int = 1
	fmt.Scan(&n)
	for i := n; i > 1; i-- {
		ans = ans * i
	}
	fmt.Println("factorial of the entered no. is=", ans)
}
