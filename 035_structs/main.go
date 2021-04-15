package main

import (
	"fmt"
	"strings"
)

func draw_line() {
	fmt.Println(strings.Repeat("*", 80))
}

func main() {

	type employee struct {
		num  int
		name string
		age  int
		ph   string
	}

	emp := employee{
		name: "Abhishek",
		num:  1,
		age:  35,
		ph:   "7799126262",
	}
	fmt.Println(emp)
	fmt.Println(emp.name)

}
