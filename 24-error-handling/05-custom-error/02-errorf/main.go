package main

import (
	"log"
	"fmt"
)

func main()  {

	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}


}

func Sqrt(v float64) (float64, error)  {
	if v < 0 {
		// %v is the default format for a value - see fmt docs
		return 0, fmt.Errorf("squareroot of negative number: %v", v)
	}

	return 42, nil

}
