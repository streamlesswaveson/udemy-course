package main

import (
	"fmt"
	"log"
)

type FooError struct {
	lat, lon string
	err      error
}

// implementing the Error() interface
func (n *FooError) Error() string {
	return fmt.Sprintf("foo error occurred: %v %v %v", n.lat, n.lon, n.err)
}

func main() {

	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(v float64) (float64, error) {
	if v < 0 {
		myerr := fmt.Errorf("squareroot of a negative number: %v", v)
		return 0, &FooError{"0.0 N,", "0.0 W", myerr}
	}

	return 42, nil

}
