package main

import (
	"fmt"
	"reflect"
	"os"
)
// Go is static typed language
// GO package level variable should be declared outside /  before main

// use `var` keyword if declaring at package level
// var name should start with '_' or any 'alphabet'
// go will always init zero value

var (
	name, course string
	module	float64
)

const speedOfLight = 18600

func pass_by_value(age int) int{
	age = age +1
	fmt.Println("var *age* in func call", age)
	return age
}

func pass_by_reference(age *int) int{
	*age = *age +1
	fmt.Println("var *age* in func call", *age)
	return *age
}

func accessenv_vars(){
	// prints map key value pairs
	fmt.Println(os.Environ())

	//prints Hosttype
	fmt.Println(os.Getenv("USER"))

	// print in loop
	// for num, val := range os.Environ(){
	// 	fmt.Println("key=",num, "  Val=", val)
	// }
}

func package_level_vars(){
	fmt.Println("Var name is set to   ", name, "  and is of type ", reflect.TypeOf(name))
	fmt.Println("Var course is set to ", course, "and is of type ", reflect.TypeOf(course))
	fmt.Println("Var module is set to ", module, "and is of type ", reflect.TypeOf(module))
	fmt.Println("const speedOfLight is set to ", speedOfLight, "and is of type ", reflect.TypeOf(speedOfLight))
	fmt.Println("**************************\n\n\n")
}

func short_dec(){
	name := "Name"
	age  := 35
	temp := 98.3
	ptr := &age

	fmt.Println("Var name is set to ", name, "and is of type ", reflect.TypeOf(name))
	fmt.Println("Var age is set to  ", age, " and is of type ", reflect.TypeOf(age))
	fmt.Println("Var temp is set to ", temp, "and is of type ", reflect.TypeOf(temp))
	fmt.Println("*memory* address of age is ", &age)
	fmt.Println("*memory* address of ptr is ", ptr, "and the value is ", *ptr)
	fmt.Println("**************************\n\n\n")

	// pass by value
	fmt.Println("var *age* before func call", age)
	pass_by_value(age)
	fmt.Println("var *age* after func call", age)
	fmt.Println("**************************\n\n\n")

	// Pass by reference
	fmt.Println("var *age* before func call", age)
	pass_by_reference(&age)
	fmt.Println("var *age* after func call", age)
	fmt.Println("**************************\n\n\n")

}

func main (){
	// Print all package levels vars
	package_level_vars()

	// Short declaration
	short_dec()

	// access environment_vars
	accessenv_vars()

}