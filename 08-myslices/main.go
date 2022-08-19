package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to Slices")

	var fruitList = []string{"Apple", "Guava", "Peach"}
	fmt.Printf("Type of Fruit list is %T\n", fruitList)

	fmt.Println(fruitList)

	fruitList = append(fruitList, "mango", "banana")

	fmt.Println("Fruilt list afer append: ", fruitList)

	for indx, val := range fruitList {
		fmt.Println(indx, val)
	}

	// 2nd way to make slices

	var vegList = make([]string, 2)

	vegList[0] = "tomato"
	vegList[1] = "potato"

	// since it lenth of 2 we cant add more but we can append

	vegList = append(vegList, "brinjal", "onion")
	fmt.Println(vegList)

	// slicing operation
	fmt.Println(vegList[:3])
	fmt.Println(vegList[1:])
	sort.Strings(vegList)
	fmt.Println(vegList)

	// how to remove value from slice based on idex

}
