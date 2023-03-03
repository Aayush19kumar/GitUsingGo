package main

import (
	"fmt"
	"time"
)

func main() {
	go greet("hello")
	greet("world")
}
func greet(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}
