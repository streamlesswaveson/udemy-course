package  main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	First string
	Last string
	Age int
}

func main()  {

	var p1 Person

	fmt.Println(p1.First, p1.Last, p1.Age)

	// note the backticks
	bs := []byte(`{"First":"Abraham","Last":"Tireshuffle","Age":24}`)

	json.Unmarshal(bs, &p1)

	fmt.Println(p1.First, p1.Last, p1.Age)

}
