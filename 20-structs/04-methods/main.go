package main

import "fmt"

type person struct {
	first string
	last string
}

// (p person) is the 'receiver'
// 'self' and 'this' are not idiomatic go
func (p person) getFullName() string  {
	return fmt.Sprintf("%s %s", p.first, p.last)
}

func main()  {

	p1 := person{"Jack", "Johnson"}

	fmt.Println(p1.getFullName())

}
