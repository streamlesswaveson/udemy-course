package main

import "fmt"

func main()  {

	myslice := []string{"Mon", "Tues", "Weds", "Thurs", "Fri", "Sat", "Sun"}

	// remove wednesday
	newslice := append(myslice[0:2], myslice[3:]...)

	fmt.Println(newslice)

}
