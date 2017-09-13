package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	<-done
	<-done
	close(c)

	// to unblock above
	// we need to take values off of chan c
	// but we never get here because we're blocked above
	// running this program will produce a deadlocked error: fatal error: all goroutines are asleep - deadlock!
	for n := range c {
		fmt.Println(n)
	}

}
