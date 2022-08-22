package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jabhishek87/mongoapp/router"
)

func main() {
	fmt.Println("Welcome to Golang and Mongo APP")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
}
