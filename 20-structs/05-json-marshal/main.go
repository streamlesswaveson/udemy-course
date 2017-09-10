package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last string
	Age int
	notExported string
}
func main()  {


	p1 := Person{"Abraham", "Tireshuffle", 24, "Cyber war meditations"}

	out, _ := json.Marshal(p1)

	// byte array
	fmt.Println(out)

	fmt.Printf("%T \n", out)

	// json output
	fmt.Println(string(out))
}
