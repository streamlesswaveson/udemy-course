package main

import "fmt"

func main()  {

	c := make(chan int)

	go func() {
		for i :=0; i< 10; i++ {
			c <- i
		}
		close(c)
	}()

	// this code blocks - waiting to receive from the channel
	for n := range c {
		fmt.Println(n)
	}


}
