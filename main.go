package main

import (
	"./learngo"
)

// "func main" is special.
//
// Go has to know where to start
//
// func main creates a starting point for Go
//
// After compiling the code,
// Go runtime will first run this function

func main() {
	// Todo make it dynamic for cli
	learngo.Ex001()
	//learngo.Statements()
	//learngo.VariableDatatypes()

	// go run main.go  abhishek shilpa
	// learngo.Greet(os.Args)
}
