package main

import "fmt"

func main() {

	fmt.Println("Welcome yo Golang Maps")

	myMaps := make(map[string]string)
	myMaps["js"] = "Javascript"
	myMaps["rb"] = "Ruby"
	myMaps["py"] = "python"
	myMaps["go"] = "golang"

	fmt.Println("Maps:", myMaps)
	fmt.Println("JS is :", myMaps["js"])

	for key, val := range myMaps {
		fmt.Println(key, "is :", val)
	}

	delete(myMaps, "rb")
	fmt.Println("Maps:", myMaps)

}
