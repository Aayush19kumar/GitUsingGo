package main

import "fmt"

func main() {

	fmt.Print("Enter Piasa: ")
	var n int
	fmt.Scan(&n)

	rs := n / 100
	ps := n % 100
	fmt.Println(rs, ".", ps)

}
