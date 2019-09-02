package main

import (
	"fmt"
	"runtime"
)



func main() {
	fmt.Println("**********")
	fmt.Println("Go Version is ", runtime.Version())
	fmt.Println("**********")

	fmt.Println("Hello Gopher")
	fmt.Println(runtime.NumCPU())

	if runtime.NumCPU() > 4 {
		fmt.Println(runtime.NumCPU(), " is greater than 4")
	}
}
