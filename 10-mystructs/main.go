package main

import "fmt"

type User struct {
	Name   string
	age    int
	email  string
	status bool
}

func (u User) SetEmail() {
	u.email = "test@go.dev"
	fmt.Println(u.email)
}

func (u *User) SetOrgEmail() {
	u.email = "ORG@go.dev"
	fmt.Println(u.email)
}

func main() {
	fmt.Println("Structs in Golang ")
	objUser := User{"Abhi", 32, "abhi@go.dev", false}
	fmt.Println(objUser)
	fmt.Printf("Details for Sruct ObjUser %+v\n", objUser)

	objUser.SetEmail()
	fmt.Printf("Details for Sruct ObjUser %+v\n", objUser)

	objUser.SetOrgEmail()
	fmt.Printf("Details for Sruct ObjUser %+v\n", objUser)
}
