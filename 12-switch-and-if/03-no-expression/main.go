package main

import "fmt"

func main()  {
	var name string;

	fmt.Scan(&name)

	switch  {

	case len(name) == 2:
		fmt.Println("Size two!")
	case name == "al": // will never run since the first case matched
		fmt.Println("I will never run")
	case name == "chad" || name == "heidi":
		fmt.Println("Rulers of the House Blair")

	default:
		fmt.Println("dunno")

	}

}
