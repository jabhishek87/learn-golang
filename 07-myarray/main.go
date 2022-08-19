package main

import "fmt"

func main() {
	fmt.Println("Array Example")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Guava"
	fruitList[3] = "Peach"

	// Please look that there is two space in between guava and Peach
	fmt.Println("Fruits in the list are : ", fruitList)
	fmt.Println("Length of fruit array : ", len(fruitList))

	var vegList = [3]string{"tomato", "potato", "onion"}
	fmt.Println("Veg in the list are : ", vegList)
	fmt.Println("Length of veg array : ", len(vegList))

	// loops

	for indx, val := range fruitList {
		fmt.Println(indx, val)
	}

}
