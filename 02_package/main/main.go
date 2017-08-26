package main

import (
	"fmt"
	"github.com/streamlesswaveson/udemy-course/02_package/stringutil"
)

func main() {

	fmt.Println(stringutil.Reverse("Hello world!"))
	fmt.Println(stringutil.MyName)

	// won't work (or compile)
	//fmt.Println(stringutil.yourName)

}
