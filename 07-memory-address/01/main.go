package main

import "fmt"

func main()  {

	foo := 42

	fmt.Println(" foo = ", foo)
	// ampersand - dereferencing the variable
	fmt.Println("foo = ", &foo)

	// e.g.
	//foo =  42
	// foo =  0xc42000e258


}
