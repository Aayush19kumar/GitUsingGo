package main

import (
	"fmt"
	"sync"
)

func main() {

	// mych := make(chan int)
	mych := make(chan int, 5) //this is a bufferd channel
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {

		val, isChannelOpen := <-mych //to check my channel is open or not
		fmt.Println(val, ":", isChannelOpen)

		// fmt.Println(<-mych)
		wg.Done()
	}(mych, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		mych <- 5
		close(mych) //we can also close the channel
		wg.Done()
	}(mych, wg)
	wg.Wait()
}

// channel can be pass in func argument (Direction is important)
