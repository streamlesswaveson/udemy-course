package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type PeopleByAge []Person

func (p PeopleByAge) Len() int {
	return len(p)
}

func (p PeopleByAge) Less(i, j int) bool {
	return p[i].age < p[j].age
}

func (p PeopleByAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Person) String() string {
	return fmt.Sprintf("%s=%d", p.name, p.age)

}

func main() {
	p1 := Person{"Chad", 43}
	p2 := Person{"Heidi", 40}
	p3 := Person{"Aidan", 13}
	p4 := Person{"Iain", 11}
	p5 := Person{"Drew", 8}

	byAge := PeopleByAge{p2, p3, p1, p4, p5}

	sort.Sort(byAge)

	fmt.Println(byAge)

}
