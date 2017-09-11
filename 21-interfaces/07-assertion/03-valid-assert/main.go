package main

import "fmt"

func main()  {

	var val interface{} = 42

	// tada! - the syntax asserts it is a particular type
	fmt.Println(val.(int) + 20)

}
