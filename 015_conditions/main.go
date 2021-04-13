package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"os"
)


func if_conditions(x_age, y_age int) {

	fmt.Println("X age=", x_age, " Y age=", y_age)

	if x_age > y_age {
		fmt.Println("X is older than Y")
	} else if y_age > x_age {
		fmt.Println("Y is older than X")
	} else {
		fmt.Println("either both are equal or something went wrong")
	}
}

func switch_stmnt(){
	// start := time.Now()
	// rand.Seed(start.Unix())

	// opts  := strings.Split("option1,option2,option3", ",")
	// sel_opt := opts[rand.Intn(len(opts))]
	//fmt.Println("Gonna work from home...", opts[rand.Intn(len(opts))])

	os_type := runtime.GOOS

	switch  os_type {
	case "linux":
		fmt.Println("OS type Linux")
	case "windows":
		fmt.Println("OS type windows")
	case "darwin":
		fmt.Println("OS type MacOS")
	default:
		fmt.Println("OS type UNKNOWN")
	}

}

func err_handling() {
	vim, vim_err := os.Open("/etc/.vimrc")
	ver, ver_err := os.Open("/etc/lsb-release")

	if vim_err != nil {
		fmt.Println("vimrc not found")
	} else {
		fmt.Println("content= ", vim)
	}

	if ver_err != nil {
		fmt.Println("lsb-release not found")
	} else {
		fmt.Println("content= ", ver)
	}
}

func main() {
	//
	if_conditions(rand.Intn(50), rand.Intn(60))

	// switch statement demo
	switch_stmnt()

	//error hadling
	err_handling()
}