package main

import "fmt"

func main()  {

	x := 'a'

	// prints the rune value - 97
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// because rune is an alias for int32
	fmt.Printf("%T\n", rune(x))

}
