package main

import (
	"strconv"
	"fmt"
)

func main()  {

	// Atoi - Ascii to Int
	s, _ := strconv.Atoi("42")
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// Int to Ascii
	o := strconv.Itoa(42)
	fmt.Printf("%T\n", o)

}
