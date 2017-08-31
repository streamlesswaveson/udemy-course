package main

import (
	"math"
	"fmt"
)

func main()  {

	var great = greatest(10,20,30,40)
	fmt.Println(great)

	great = greatest()
	fmt.Println(great)

	slice := []int{1,2,3,4}
	great = greatest(slice...)
	fmt.Println(great)

}

func greatest(nums ... int) int  {
	var biggest int = math.MinInt64

	for _, thing := range nums {
		if thing > biggest {
			biggest = thing
		}
	}
	return biggest

}
