package main

import "fmt"

func main() {
	fmt.Println("enter any no. to check whether it's palindrome or not")
	var n int
	fmt.Scan(&n)
	var storing = n
	var ans = 0
	for n != 0 {
		r := n % 10
		ans = ans*10 + r
		n = n / 10
	}

	if storing == ans {
		fmt.Print("no. is palindrome\n")
	} else {
		fmt.Println("no. is not palindrome")
	}

}
