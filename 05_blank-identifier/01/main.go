package main

import (
	http "net/http"
	"io/ioutil"
	"fmt"
)

// blank identifer (_) tells golang that you don't need the result value
func main()  {


	resp, _ := http.Get("https://www.google.com")

	page, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	fmt.Printf("%s", page)

}
