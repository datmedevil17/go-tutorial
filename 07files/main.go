package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to Files")
	content := "This will be going in the file."

	file, err := os.Create("./mylogfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	check(err)

	fmt.Println("Length of content written to file:", length)

	defer file.Close()
	readFile("./mylogfile.txt")
}


func readFile(filename string){
	databyte, err := os.ReadFile(filename)
check(err)
	fmt.Println("Contents of the file are:")
	fmt.Println(string(databyte))
}

func check(err error){
	if err != nil {
		panic(err)
	}
}