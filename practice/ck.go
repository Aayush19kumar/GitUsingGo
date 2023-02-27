package main

import (
	"errors"
	"fmt"
)

func main() {

	// fmt.Println(show("aayush"))
	fmt.Println(show2("testing"))
}

func show(s string) (string, error) {

	return s, errors.New("no error")
}

func show2(s string) (string, error) {
	return s, fmt.Errorf("not")
}
