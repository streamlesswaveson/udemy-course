package main

import "fmt"

func main()  {

	var mymap = map[string]string{}

	fmt.Println("is nil? ", mymap == nil)

	mymap["dinner"] = "Chicken"

	fmt.Println( "len=", len(mymap))
}
