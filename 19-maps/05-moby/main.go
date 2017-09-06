package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)

func main()  {

	response, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(response.Body)
	defer response.Body.Close()

	scanner.Split(bufio.ScanWords)

	buckets := make([]int, 200)

	for scanner.Scan() {
		n := moduloBucket(scanner.Text(), 12)
		buckets[n]++
	}

	fmt.Println(buckets[0:12])
	//fmt.Println(buckets[65:123])

}

func hashBucket(s string) int {

	return int(s[0])
}

func moduloBucket(s string, bucket int) int  {

	letter := hashBucket(s)

	return letter % bucket

}
