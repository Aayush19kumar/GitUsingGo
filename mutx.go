package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var m sync.Mutex
var storing []string

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
	fmt.Println(storing)

}

func StatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("oops in endpoint")
	} else {
		m.Lock()
		storing = append(storing, endpoint)
		m.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
