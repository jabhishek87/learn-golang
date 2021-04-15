package main

import (
	"fmt"
	"strings"
)

func draw_line() {
	fmt.Println(strings.Repeat("*", 80))
}

func _maps_init() {

	// declaring maps using make
	// ictionary with name , quantity
	inventory := make(map[string]int)
	inventory["mango"] = 30
	inventory["orange"] = 10

	// declaring maps in single line short hand

	new_inv := map[string]int{
		"guava":   50,
		"avacado": 10,
		"kiwi":    5,
	}

	fmt.Printf("Inventories:\n %v, \n %v \n", inventory, new_inv)
	draw_line()
}

func _maps_manupulating() {
	// looping through maps, it will different every time we run as it is
	// unordered
	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}

	for key, val := range testMap {
		fmt.Println("Key:", key, " val:", val)
	}
	draw_line()

	// print the value of known key
	fmt.Println(testMap["C"])

	// updatinh value
	testMap["C"] = 02

	// addng new key val
	testMap["F"] = 6

	delete(testMap, "A")

	fmt.Println(testMap)
}

func main() {

	_maps_init()

	_maps_manupulating()

}
