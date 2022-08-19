package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://dummy.restapiexample.com/"

type sResponse struct {
	Status string `json:"status"`
	Data   []struct {
		ID      int    `json:"id"`
		EmpName string `json:"employee_name"`
		EmpSal  int    `json:"employee_salary"`
		EmpAge  int    `json:"employee_age"`
		EmpImg  string `json:"profile_image"`
	} `json:"data"`
	Message string `json:"message"`
}

func main() {
	fmt.Println("Golang Web Request")

	// calling a url which gives Html as response
	// getHtml(url)
	getJson(url + "api/v1/employees")
	//fmt.Println(content)
}

func getUrl(url string) *http.Response {
	resp, err := http.Get(url)
	checkErr(err)

	fmt.Printf("Type of resp : %T\n", resp)
	// fmt.Printf("resp : %+v\n", resp)

	return resp
}

func getJson(url string) {
	resp := getUrl(url)
	defer resp.Body.Close() // caller's responsible to close
	fmt.Println("Response Code: ", resp.StatusCode)
	fmt.Println("Response Content Type : ", resp.Header.Get("Content-type"))

	jsonStr, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	var result sResponse
	if err := json.Unmarshal(jsonStr, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	// fmt.Printf("result : %+v\n", result)

	for _, rec := range result.Data {
		fmt.Println(rec.EmpName)
	}
}

func getHtml(url string) {
	resp := getUrl(url)
	defer resp.Body.Close() // caller's responsible to close
	// Since Response is Databytes we need to converting to string
	dataBytes, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(dataBytes))

	fmt.Println("Response Code: ", resp.StatusCode)
	fmt.Println("Response Content Type : ", resp.Header.Get("Content-type"))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
