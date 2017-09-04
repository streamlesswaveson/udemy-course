package main

import "fmt"

func main()  {

	stuff := []string{"a","B","c"}

	fmt.Println(stuff)

	fmt.Println(cap(stuff))

	stuff = append(stuff, "Something else")

	fmt.Println(stuff)

	fmt.Println(cap(stuff))

	fmt.Println(stuff[len(stuff) - 1])

}
