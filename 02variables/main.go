package main

import "fmt"

const LoginToken string = "abc123"

func main() {
	var username string = "Rakshit"
	fmt.Println(username)
	fmt.Printf("Type of username variable is %T\n", username)
     
		var variable bool 
	fmt.Println(variable)
	fmt.Printf("Type of username variable is %T\n", variable)

	var website = "rakshit.com"
	fmt.Println(website)
//not used in Global
	num :=20
	fmt.Println(num)

	fmt.Println("Login token is:", LoginToken)

}
