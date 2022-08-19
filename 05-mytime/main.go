package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time Module")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006"))

	createDate := time.Date(1987, time.April, 9, 23, 23, 23, 0, time.Local)
	fmt.Println(createDate.Format("02-01-2006, Monday"))

}
