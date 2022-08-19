package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers ")

	// Declaring Pointers
	var ptr *int
	fmt.Println("value of Pointer ", ptr)

	// Assigning Memory addres to Pointer

	myNumber := 23
	var myptr = &myNumber

	fmt.Println("Memory Address for Mynumber ", myptr)
	fmt.Println("value for Mynumber ", *myptr)

	*myptr = *myptr * 2
	fmt.Println("New value of myNumber ", myNumber)
}
