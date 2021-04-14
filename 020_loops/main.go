package main
/*
For loops in go
a) Format
	for <expr> {
		<code>
	}
b) Infinite loop
	for {
		fmt.Println("Name")
	}
*/

import (
	"fmt"
	"time"
)
func helper(var_t int) {

	fmt.Println(var_t)
	time.Sleep(1* time.Second)

	if var_t == 1 {
		fmt.Println("Bye Bye !!!")
	}
}
func countdown_timer(timer int){
	for var_t := timer; var_t > 0; var_t-- {
		helper(var_t)

	}
	fmt.Println("END countdown_timer *******************")
}

func for_while(timer int) {
	for timer > 0 {
		helper(timer)
		timer--
	}
	fmt.Println("END for_while *******************")
}

func for_range(elements []string) {
	fmt.Println("Nicknames are :")
	for k, i := range elements {
		fmt.Println("\t",k,i)
	}
}

func _continue(){
	// This will only print odd number
	for i := 10; i >=0; i-- {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		time.Sleep(2* time.Second)
	}
}

func _break() {

	for i := 10; i >=0; i-- {
		// this will not print anything after 5
		if i == 5 {
			break;
		}
		fmt.Println(i)
	}
}

func main(){
	// simple for loop example
	countdown_timer(10)

	// wile loop example
	for_while(5)

	// looping through range example
	slice_name := []string{"Abhishek", "sonu", "ABH", "ABHI"}
	for_range(slice_name)

	// continue example
	_continue()

	// break example
	_break()

}