package main

import (
	"sync"
	"fmt"
	"time"
)

var wg sync.WaitGroup

func main()  {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	
}

func foo()  {

	for i := 0; i< 50; i++ {
		fmt.Printf("Foo %d\n", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}

	wg.Done()

}

func bar()  {
	for i := 0; i< 50; i++ {
		fmt.Printf("Bar %d\n", i)
		time.Sleep(time.Duration(12 * time.Millisecond))
	}

	wg.Done()

}
