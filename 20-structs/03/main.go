package main

import "fmt"

type foo int

func main()  {

	var bar foo
	var blerg int

	bar = 23
	blerg = 17

	// won't work because they are different types
	//fmt.Println(bar + blerg)

	// will work after converting bar to an int
	fmt.Println(int(bar) + blerg)

}
