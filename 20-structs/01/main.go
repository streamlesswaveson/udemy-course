package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main()  {

	p1 := person{"James", "Bond", 30}
	p2 := person{"Money", "Penny", 28}

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

}
