package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("My Time Module")

	presentTime := time.Now()
	fmt.Println("Present time is:", presentTime.Format("02-Jan-2006 15:04:05 Monday"))

	
}
