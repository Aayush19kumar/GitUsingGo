package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func world(fn func()) {

	fn()
}

func main() {

	world(hello)
}
