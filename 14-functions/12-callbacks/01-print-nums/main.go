package _1_print_nums

import "fmt"

func main()  {

	visit([]int{1,2,3,4}, func(i int) {
		fmt.Println("visiting", i)
	})

}

func visit(numbers []int, callback func(int))  {

	for _, thing := range numbers {
		callback(thing)
	}
}
