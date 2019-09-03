package main

import (
	"fmt"
	"os"
)

func printArgs(args []string) {
	fmt.Printf("%#v\n", args)
}

func greeter(args []string) {
	fmt.Println("I have to Greet", len(args)-1, "People")
	fmt.Printf("\n\n\n")

	for i := 1; i < len(args); i++ {
		name := args[i]
		fmt.Println("Hello, ", name)
	}
}

// build and run
// go build -o src/004_project_greeter/greeter src/004_project_greeter/*.go
// ./src/004_project_greeter/greeter
func main() {
	printArgs(os.Args)
	greeter(os.Args)
}
