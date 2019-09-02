package main

import "fmt"

// package scope
func hey() {
	//block scope
	fmt.Println("Hey Gopher ! from hey.go file")
}
