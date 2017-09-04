package main

import "fmt"

func main()  {

	var things [256]int

	for i :=0; i<256; i++ {
		things[i] = i
	}

	for _, v := range things {
		fmt.Printf("%v - %T - %b\n", v, v, v)
	}

}
