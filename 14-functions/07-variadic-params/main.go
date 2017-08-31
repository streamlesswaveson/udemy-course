package main

import "fmt"

func main()  {

	fmt.Println(average(1,2,3,4,5, 6, 7, 8, 9, 10))

}

func average(nums ... float64) float64  {

	var total float64
	// the blank identifier is the index
	for _, thing := range nums {
		total += thing
	}

	return total / float64(len(nums))

}
