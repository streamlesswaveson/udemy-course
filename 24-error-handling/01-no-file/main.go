package main

import (
	"os"
	"fmt"
	"log"
)

func main()  {

	_, err := os.Open("no-file.txt")

	if err != nil {
		fmt.Println("err happened", err)
		log.Println("logged error", err)
	}

}
