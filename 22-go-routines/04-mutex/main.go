package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var wg sync.WaitGroup

var counter int

var mutex sync.Mutex

func main()  {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Printf("Final counter: %d\n", counter)

}

func incrementor(name string)  {
	for i := 0; i< 20; i++ {
		mutex.Lock()
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		counter = x
		fmt.Println(name, i, " Counter: ", counter)
		mutex.Unlock()
	}
	wg.Done()

}
