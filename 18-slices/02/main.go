package main

import "fmt"

// appending a slice to a slice
func main()  {

	slice1 := []int{1,2,3,4}

	slice2 := []int{9,8,7,6}

	// note the variadic syntax
	merged := append(slice1, slice2...)

	fmt.Println(merged)

}
