package main

import "fmt"

type person struct {
	name string
	age  int
	job  job
}

type job struct {
	name   string
	salary float32
}

func main() {

	p1 := person{
		name: "Chad",
		age:  43,
		job: job{
			name:   "coding guy",
			salary: 1.23,
		},
	}

	fmt.Println(p1.job.salary)

}
