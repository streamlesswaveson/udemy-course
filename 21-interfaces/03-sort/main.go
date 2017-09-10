package main

import (
	"sort"
	"fmt"
)

type people  []string

func main()  {

	// 1

	peeps := people{"Zeno", "John", "Al", "Jenny"}

	sort.Strings(peeps)

	fmt.Println(peeps)
	sort.Sort(sort.Reverse(sort.StringSlice(peeps)))
	fmt.Println(peeps)

	// 2
	p := []string{"Zeno", "John", "Al", "Jenny"}

	sort.Strings(p)

	fmt.Println(p)

	// reversed
	sort.Sort(sort.Reverse(sort.StringSlice(p)))

	fmt.Println(p)

	/// 3

	i := []int{7, 4, 8, 2, 9, 19, 12, 32, 3 }

	sort.Ints(i)
	fmt.Println(i)

	sort.Sort(sort.Reverse(sort.IntSlice(i)))

	fmt.Println(i)


	fmt.Println("ape" < "apple")

}
