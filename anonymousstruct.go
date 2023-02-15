package main

import "fmt"

func main() {
	st := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  23,
	}
	fmt.Print(st)
}
