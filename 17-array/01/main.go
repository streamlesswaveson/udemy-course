package main

import "fmt"

func main()  {

	var things [58]string

	// prints an array of 58 empty strings
	fmt.Println(things)

	fmt.Println(things[57])

	// interesting - you'll get an compiler warning if you try to do this
	//fmt.Println(things[58]) -- try to print the non-existent 59th element

	for i :=65; i <=122; i++ {
		things[i-65] = string(i)
		
	}

	fmt.Println(things)

	fmt.Println(things[57])


}
