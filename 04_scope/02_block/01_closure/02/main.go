package main

import "fmt"

// like other functional languages you can create function factories
// note the return type:  func() int
func wrapper() func() int  {

	x := 0
	return func() int {
		x++
		return x
	}

}
func main()  {

	hey := wrapper()
	jude := wrapper()

	fmt.Println(hey())
	fmt.Println(hey())
	fmt.Println(hey())

	fmt.Println(jude())

}
