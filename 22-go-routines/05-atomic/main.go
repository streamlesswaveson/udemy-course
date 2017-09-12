package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
	"sync/atomic"
)

var wg sync.WaitGroup

var counter int64

func main()  {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Printf("Final counter: %d\n", counter)

}

func incrementor(name string)  {
	for i := 0; i< 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)

		atomic.AddInt64(&counter, 1)

		fmt.Println(name, i, " Counter: ", counter)
	}
	wg.Done()

}
