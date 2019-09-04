package learngo

import (
	"fmt"
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

// Greet the passed args
func Greet(args []string) {
	printArgs(args)
	greeter(args)
}
