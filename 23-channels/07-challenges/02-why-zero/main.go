package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}



}

//func main()  {
//
//	c := make(chan int)
//
//	go func() {
//		for i :=0; i<10; i++ {
//			c <- i
//		}
//	}()
//
//	fmt.Println(<-c)
//
//
//
//}