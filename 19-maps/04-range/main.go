package main

import "fmt"

func main()  {

	mymap := map[string]string {
		"Dad":"Chad",
		"Mom": "Heidi",
		"Kids": "Aidan, Iain, Andrew",
	}

	for k, xiv := range  mymap {
		fmt.Println("key=", k, "value=", xiv)
	}

}
