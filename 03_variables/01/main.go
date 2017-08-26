package main

import "fmt"

// declaring variable with shorthand
// note - the type is not declared
func main()  {
	a := 10
	b := "golang"
	c := 4.17
	d := true

	// %v is magic - it uses the default format for the variable type
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n",d )

	// %T prints the type
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n",d )
}
