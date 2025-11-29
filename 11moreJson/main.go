package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Platform string `json:"platform"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	DecodeJson()
}
func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"Angular Bootcamp", 199, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"VueJS Bootcamp", 299, "learncodeonline.in", "abc123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	check(err)

	fmt.Printf("%s\n", finalJson)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}


func DecodeJson(){
	jsonDataFromWeb := []byte(`[
		{"name": "ReactJS Bootcamp", "price": 299, "platform": "learncodeonline.in", "password": "abc123", "tags": ["web-dev", "js"]},
		{"name": "Angular Bootcamp", "price": 199, "platform": "learncodeonline.in", "password": "abc123", "tags": ["web-dev", "js"]},
		{"name": "VueJS Bootcamp", "price": 299, "platform": "learncodeonline.in", "password": "abc123", "tags": null}
	]`)

	var courses []course
	err := json.Unmarshal(jsonDataFromWeb, &courses)
	check(err)

	for _, course := range courses {
		fmt.Printf("Name: %s, Price: %d, Platform: %s, Tags: %v\n", course.Name, course.Price, course.Platform, course.Tags)
	}

	var myData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myData)
	fmt.Println(myData)

	for _, course := range myData {
		courseMap := course.(map[string]interface{})
		fmt.Printf("Name: %s, Price: %d, Platform: %s, Tags: %v\n",
			courseMap["name"], courseMap["price"], courseMap["platform"], courseMap["tags"])
	}
}