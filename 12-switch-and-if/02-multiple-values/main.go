package main

import "fmt"

func main() {

	var name string
	fmt.Scan(&name)

	// multiple evaluations allowed
	switch name {

	case "chad", "heidi":
		fmt.Println("the parents!")
	case "aidan", "iain", "andrew":
		fmt.Println("the kids!")
	default:
		fmt.Println("shrug")

	}

}
