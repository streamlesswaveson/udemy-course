package  main

import "fmt"

func main()  {

	var foo interface{} = 42

	fmt.Printf("foo is %T\n", foo)
	// but... won't work
	fmt.Println(foo + 2)

}
