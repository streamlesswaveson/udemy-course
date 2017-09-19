package main

import "fmt"

func main() {

	fmt.Println(factorial(4))

	fmt.Println(factgo(4))

}

func factgo(num int) int {

	c := make(chan int)

	go func() {
		for i := 1; i <= num; i++ {
			c <- i
		}
		close(c)

	}()

	total := 1
	for i := range c {
		total *= i
	}
	return total

}

func factorial(num int) int {
	total := 1
	for i := 1; i <= num; i++ {
		total *= i
	}
	return total

}
