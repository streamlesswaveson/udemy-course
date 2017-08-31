package main

import "fmt"

func main()  {

	out := filter([]int{1,2,3,4,5, 6, 7, 8, 9, 10}, func(i int) bool {
		return i % 2 == 0
	})

	fmt.Println(out)
}

func filter(nums []int, callback func(int) bool) []int {

	var out []int;

	for _, thing := range nums {
		if callback(thing) {
			out = append(out, thing)
		}
	}
	return out

}
