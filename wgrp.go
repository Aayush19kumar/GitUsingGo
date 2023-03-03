package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //in most of the cases it's of pointer type and we'll pass a reference of it

func main() {

	website := []string{

		"http://www.google.com",
		"http://www.fb.com",
		"http://www.linkedin.com",
	}
	for _, v := range website {
		go StatusCode(v)
		wg.Add(1)

	}
	wg.Wait()

}

func StatusCode(s string) {
	defer wg.Done()
	res, err := http.Get(s)
	if err != nil {
		fmt.Println("oops in endpoint")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, s)
	}
}
