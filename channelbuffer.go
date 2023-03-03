package main

import "fmt"

func main() {

	m := make(chan string, 2)
	m <- "Hello"
	m <- "World"
	fmt.Println(<-m, <-m)

}
