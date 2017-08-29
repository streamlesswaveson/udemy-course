package main

import "fmt"

func main()  {

	for i := 5000; i<=5100; i++ {

		// displaying some three byte characters
		fmt.Println(i, " - ", string(i), " - ",  []byte(string(i)))
	}

}
