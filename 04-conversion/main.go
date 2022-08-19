package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welocme to COnversion")

	fmt.Println("Enter the Rating on scale of 1 - 5")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}
	iRating, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		panic(err)
	}

	fmt.Println("Added 1 to your rating: ", iRating+1)
}
