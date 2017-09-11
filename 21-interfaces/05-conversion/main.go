package main

import "fmt"

func main()  {

	var a int = 23
	var b float64 = 23.4566

	fmt.Println(int(b) + a)

	var c rune = 'a'
	var d int32 = 'b'

	fmt.Println(c)
	fmt.Println(string(c))

	fmt.Println(d)
	fmt.Println(string(d))

}
