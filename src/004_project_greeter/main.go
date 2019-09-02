package main

import (
	"fmt"
	"os"
)

func print_args(args []string) {
	fmt.Printf("%#v\n", args)
}

func greeter(args []string) {
	for i := 1; i < len(args); i++ {
		name := args[i]
		fmt.Println("Hello, ", name)
	}
}

// build and run
// go build -o src/004_project_greeter/greeter src/004_project_greeter/*.go
// ./src/004_project_greeter/greeter
func main() {
	print_args(os.Args)
	greeter(os.Args)
}
