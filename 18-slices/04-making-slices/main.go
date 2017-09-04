package main

import "fmt"

func main()  {

	people := make([]string, 2)

	fmt.Printf("people nil? %t \n", people == nil)

	people[0] = "Chad"
	people[1] = "Heidi"
	// index out of range
	//people[2] = "Aidan"

	people = append(people, "Aidan")

	fmt.Println(people)

	fmt.Println(cap(people), len(people))

	people = append(people, "Iain")

	fmt.Println(cap(people), len(people))

	people = append(people, "Andrew")

	fmt.Println(cap(people), len(people))
}
