package main

import (
	"fmt"
)

func main() {
	str := "Hello Aayush"
	ch := []rune(str)
	fmt.Println(Rev(ch))

}

func Rev(char []rune) string {
	var temp []rune
	for i := len(char) - 1; i >= 0; i-- {
		temp = append(temp, char[i])
	}
	return string(temp)
}
