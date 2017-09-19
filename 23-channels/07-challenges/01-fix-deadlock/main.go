package main

import "fmt"

func main()  {

	c := make(chan int)
	// the solution is to wrap this line in a goroutine
	go func() {
		c <- 1
	}()

	fmt.Println(<-c)
}

// this version results in deadlock
//func main()  {
//
//	c := make(chan int)
	// this line puts something on a channel and then waits - the println statement will never execute
	//c <- 1
	//fmt.Println(<-c)
//
//}
