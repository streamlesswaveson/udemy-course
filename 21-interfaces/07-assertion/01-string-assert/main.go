package _1_string_assert

import "fmt"

func main()  {

	var name interface{}  = 12

	str, ok := name.(string)
	if ok {
		fmt.Printf("It's a string %s", str)
	} else {
		fmt.Printf("Not a string. It's a %T", name)
	}

}
