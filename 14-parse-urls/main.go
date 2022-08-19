package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://emaple.com:3000/learn?name=test&sort=desc&filter=name"

func main() {
	fmt.Println("URL Parsing in Golang")
	res, _ := url.Parse(myurl)
	fmt.Println(res)
	// fmt.Printf("Type of res is %T,\n", res)
	// fmt.Printf("Res is %+v\n\n", res)
	// fmt.Printf("Res is %#v", res)

	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)
	fmt.Println(res.Path)
	fmt.Println(res.RawPath)

	// Split the params
	params := res.Query()
	fmt.Println(params)
	// fmt.Printf("params is %T\n\n", params)
	fmt.Println(params["name"])

}
