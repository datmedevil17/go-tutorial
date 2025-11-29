package main

import "fmt"

func main(){
	fmt.Println("Hello, World!")

	// var ptr *int
	// fmt.Println("Value of pointer ptr is:", ptr)
	myNumber:=23
	var ptr = &myNumber

	fmt.Println("Value of myNumber is:", myNumber)
   	fmt.Println("Address of myNumber is:", ptr)
}
