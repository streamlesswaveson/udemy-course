package main

import (
	"strings"
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last string
	Age int
	notExported int
}

func main()  {

	var p1 Person

	rdr := strings.NewReader(`{"First":"Sonny","Last":"Chickensfarm","Age":24}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1.First, p1.Last, p1.Age, p1.notExported)

}
