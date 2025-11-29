package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://dummyjson.com/posts/"

type Post struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	UserID    int      `json:"userId"`
	Tags      []string `json:"tags"`
	Reactions struct {
		Likes    int `json:"likes"`
		Dislikes int `json:"dislikes"`
	} `json:"reactions"`
}

type PostsResponse struct {
	Posts []Post `json:"posts"`
	Total int    `json:"total"`
	Skip  int    `json:"skip"`
	Limit int    `json:"limit"`
}

func main() {

	response, err := http.Get(url)
	check(err)
	// fmt.Printf("Response is of type: %T",response)

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	check(err)

	// Parse the JSON response
	var postsResponse PostsResponse
err = json.Unmarshal(body, &postsResponse)
	check(err)

	fmt.Printf("Total posts: %d\n", postsResponse.Total)
	fmt.Printf("Number of posts in response: %d\n\n", len(postsResponse.Posts))

	if len(postsResponse.Posts) > 0 {
		firstPost := postsResponse.Posts[0]
		fmt.Println("=== FIRST POST ===")
		fmt.Printf("ID: %d\n", firstPost.ID)
		fmt.Printf("Title: %s\n", firstPost.Title)
		fmt.Printf("Body: %s\n", firstPost.Body)
		fmt.Printf("User ID: %d\n", firstPost.UserID)
		fmt.Printf("Tags: %v\n", firstPost.Tags)
		fmt.Printf("Likes: %d, Dislikes: %d\n", firstPost.Reactions.Likes, firstPost.Reactions.Dislikes)
	} else {
		fmt.Println("No posts found in the response")
	}

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
