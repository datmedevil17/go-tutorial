package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcomeMessage := "Welcome to user input in golang"
	println(welcomeMessage)
	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	name,_ := reader.ReadString('\n')
	fmt.Println("Hello", name)
	fmt.Printf("Type of name is %T\n", name)
}
