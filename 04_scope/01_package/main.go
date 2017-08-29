package main

import "fmt"

// scope of x is the whole package
var x int = 42

func main() {
	fmt.Println(x)

	foo()
}

func foo()  {
	fmt.Println(x)

}
