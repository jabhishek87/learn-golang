package main

import (
	"fmt"
	"strings"
)

/*

array
  number list of same type
  fixed length

slices
  number list of same type
  dynamic length

*/
func draw_line() {
	fmt.Println(strings.Repeat("*", 80))
}

func _array() {

}

func _slice() {
	// Declares a slice called myCourses
	// syntax make([]<type>, length, capacity)
	myCourses := make([]string, 5, 10)
	fmt.Printf(
		"Length is %d, capacity is %d value %v\n", len(myCourses), cap(myCourses), myCourses)
	draw_line()

	// declare a slice without capacity
	mySlice := make([]string, 5)
	fmt.Printf(
		"Length is %d, capacity is %d value %v\n", len(mySlice), cap(mySlice), mySlice)
	draw_line()

	// declare shorthand with init
	mySlice_2 := []string{"Abhisek", "sonu", "ABH", "ABHI"}
	fmt.Printf(
		"Length is %d, capacity is %d value %v\n", len(mySlice_2), cap(mySlice_2), mySlice_2)
	draw_line()

	// slicing a slice
	// declare shorthand with init
	mySlice_3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf(
		"Length is %d, capacity is %d value %v\n", len(mySlice_3), cap(mySlice_3), mySlice_3)
	draw_line()

	// slicing
	// slice_var[<start_inclusive>:<end_exclusive>]
	fmt.Println("Original Slice. ", mySlice_3)
	fmt.Println("subslice [0:5]. ", mySlice_3[0:5])
	fmt.Println("subslice [5:8]. ", mySlice_3[5:8])
	fmt.Println("subslice [2:]. ", mySlice_3[2:])
	fmt.Println("subslice [:5]. ", mySlice_3[:5])
	draw_line()

	// slice appending without capacity
	mySlice_5 := []int{0}
	fmt.Printf("Length is %d, capacity is %d\n", len(mySlice_5), cap(mySlice_5))

	// This is the demonastration how slice increases its value
	// every time when its reach the threshold it doubled itself
	for i := 1; i < 17; i++ {
		mySlice_5 = append(mySlice_5, i)
		fmt.Printf(
			"\nCapacity is %d, length %d, value=%v",
			cap(mySlice_5), len(mySlice_5), mySlice_5)
	}
	fmt.Println()
	draw_line()

}

func main() {

	// array demo
	_array()

	// slice demo
	_slice()

}
