package main

import "fmt"

// package scope
func bye() {
	// block scope
	fmt.Println("Bye Gopher ! from bye.go file")
}
