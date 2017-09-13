package main

import "fmt"

// multiple consumers of a channel
func main()  {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i :=0; i<10; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println(n, "foo")
		}
		done <- true
	}()

	go func() {
		for n := range c {
			fmt.Println(n, "bar")
		}
		done <- true
	}()

	<- done
	<- done

}
