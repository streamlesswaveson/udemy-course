package main

import (
	"fmt"
	"time"
)

func main()  {

	c :=  make(chan int)
	// you can also bound/buffer the channel with an additional arg
	//c :=  make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			// the code stops until something takes the element off
			// the nature of the unbuffered channel
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)

}
