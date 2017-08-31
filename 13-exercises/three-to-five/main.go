package main

import "fmt"

func main()  {

	var ceiling int
	fmt.Scan(&ceiling)

	var sum int

	for i := 0; i< ceiling; i++ {
		if i % 5 == 0 {
			sum += i
		} else if i % 3 == 0 {
			sum += i
		}
	}

	fmt.Println("Sum is ", sum)
	
}
