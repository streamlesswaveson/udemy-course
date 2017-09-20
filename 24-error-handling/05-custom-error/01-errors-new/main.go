package main

import (
	"errors"
	"log"
)

func main()  {

	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}


}

func Sqrt(v float64) (float64, error)  {
	if v < 0 {
		return 0, errors.New("squareroot of a negative number")
	}

	return 42, nil

}
