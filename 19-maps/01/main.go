package main

import "fmt"

func main()  {

	mymap := make(map[string]int)

	mymap["k1"] = 1
	mymap["k2"] = 2

	fmt.Println(mymap)

	k1 := mymap["k1"]
	fmt.Println(k1)

	_, ok := mymap["k2"]
	fmt.Printf("Key exists? %t\n", ok)

	delete(mymap, "k2")

	fmt.Println(mymap)

	_, ok = mymap["k2"]

	fmt.Printf("Key exists? %t\n", ok)


}
