package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Golang Switch")

	rand.Seed(time.Now().UnixNano())

	// since in not inclusive it wil provide between 0 -  5
	// if we add 1 it wil become 1-6
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of die is :", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 now yoo can open")
	case 2, 3, 4, 5:
		fmt.Println("you can move to ", diceNumber, "Spot")

	case 6:
		fmt.Println("you can move to 6 Spot and roll dice again")
	}

}
