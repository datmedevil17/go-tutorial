package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/datmedevil17/mongo/router"
)

func main() {
	fmt.Println("Server is getting started ...")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server is listening to 4000" )

}
