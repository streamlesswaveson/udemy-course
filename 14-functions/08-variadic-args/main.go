package main

import "fmt"

func main()  {

	data := []float64{ 1,2,3,4,5, 6, 7, 8, 9, 10}

	// the ... "opens up" the slice for reading
	fmt.Println(average(data...))

}

func average(nums ... float64) float64  {

	var total float64
	// the blank identifier is the index
	for _, thing := range nums {
		total += thing
	}

	return total / float64(len(nums))

}
