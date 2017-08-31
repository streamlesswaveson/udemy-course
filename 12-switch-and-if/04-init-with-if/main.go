package main

import "fmt"

func main()  {

	// note the scope of name - only within the if block
	if name := getName(); name == "Chocolate" {
		fmt.Println(name, " is Delicious")
	}


}
func getName() interface{} {

	var name string;
	fmt.Scan(&name)
	return name
}
