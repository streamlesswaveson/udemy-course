package main

import "fmt"

func main()  {

	var foo int = 42

	fmt.Println(foo)
	fmt.Println(&foo)

	// *int is the type - a an int pointer, in this case
	var bar *int = &foo

	fmt.Println(bar)
	// now dereferenced - give me the value at that address, not the address
	fmt.Println(*bar)

	*bar = 43

	// should be the same value
	fmt.Println(foo)
	fmt.Println(*bar)
}
