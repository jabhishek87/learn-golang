package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "./logfile.out"
	fmt.Println("Handling files in Golang")

	strContent := "THis is the contents of file"

	file, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	file.WriteString(strContent)
	defer file.Close()

	fmt.Println("COntents of file : \n", readFile(fileName))

}

func readFile(filename string) string {
	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}
	return string(contents)
}
