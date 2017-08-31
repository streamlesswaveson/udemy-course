package main

import "fmt"

func main()  {

	var myval, mybool = half(1)
	fmt.Println(myval, mybool)

	var fooval, foobool = half(2)
	fmt.Println(fooval, foobool)

}

func half(yourint int) (int, bool)  {

	var halved = yourint / 2

	var isEven bool = yourint % 2 == 0

	return halved, isEven

}
