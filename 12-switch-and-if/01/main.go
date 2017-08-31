package main

import "fmt"

func main() {

	var name string

	fmt.Print("name?")
	fmt.Scan(&name)

	// fallthrough is optional! Must be made explicit
	// break is not needed as in java, etc.
	switch name {

	case "chad":
		fmt.Println("that's you!")
	case "heidi":
		fmt.Println("that's your wife!")
	case "aidan":
		fallthrough
	case "iain":
		fallthrough
	case "andrew":
		fmt.Println("that's your son!")
	default:
		fmt.Println("not sure who that is")

	}

}
