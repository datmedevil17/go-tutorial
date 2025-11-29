package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	myUrl := "http://localhost:3000/api/posts"

	resp, err := http.Get(myUrl)

	check(err)

	defer resp.Body.Close()
	fmt.Println(resp.Status, resp.StatusCode, resp.ContentLength)

	content, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(content))
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count:", byteCount)
	fmt.Println("Response string:", responseString.String())
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:3000/api/posts"

	requestBody := strings.NewReader(`
	{
		"title": "Learn Go Programming",
		"body": "A comprehensive guide to Go programming language.",
		"author": "learncodeonline.io"
	}`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	check(err)

	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	check(err)

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count:", byteCount)
	fmt.Println("Response string:", responseString.String())
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:3000/api/postform"

	data := url.Values{}
	data.Add("title", "Go Form Post")
	data.Add("body", "This is a form submission from Go")
	data.Add("author", "John Doe")
	data.Add("category", "Programming")
	data.Add("tags", "go,form,http")

	response, err := http.PostForm(myUrl, data)
	check(err)
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count:", byteCount)
	fmt.Println("Response string:", responseString.String())

}


