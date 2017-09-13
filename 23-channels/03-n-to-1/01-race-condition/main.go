package main

import (
	"sync"
	"fmt"
)

func main()  {

	c := make(chan int)

	var wg sync.WaitGroup

	go func() {
		// here's the race condition
		wg.Add(1)
		for i :=0; i< 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		// here's the race condition
		wg.Add(1)
		for i :=0; i< 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}
