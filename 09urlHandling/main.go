package main

import (
	"fmt"
	"net/url"
)

const myUrl ="https://dummyjson.com/posts"

func main(){

	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of query params is: %T\n", qparams)


}
