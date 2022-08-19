package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to User input\n"
	fmt.Printf("%s", welcome)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for Pizza on a scale of 1-5")
	input, _ := reader.ReadString('\n')

	fmt.Printf("Thanks for Rating Us with %s", input)

}
