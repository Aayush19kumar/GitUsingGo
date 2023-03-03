package main

import "fmt"

func main() {
	v := make(chan string)
	go show(v)
	v <- "display"

}

func show(s chan string) {
	fmt.Println("string is :", <-s)
}
