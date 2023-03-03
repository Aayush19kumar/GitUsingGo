package main

import (
	"fmt"
	"time"
)

func main() {
	go Greet("Hello")
	Greet("World")
}

func Greet(s string) {
	for i, _ := range []int{1, 2, 4, 5, 4} {
		fmt.Println("Hello GoRoutine", s, i)
		time.Sleep(3 * time.Second)
	}
}
