package main

// file scope
import "fmt"

// package scope
// can directely call any function defined in same package in other files
// as well becasue of package scope, i.e why  we can call hey and bye func
// defined in others file with importing it because of scope
func main() {
	// block scope
	hey()
	fmt.Println("Hello Gopher ! from main.go")
	bye()
}
