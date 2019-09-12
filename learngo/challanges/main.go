package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

//const m map[string]string

// Authenticate func to check username / password as args
func Authenticate(args []string) string {

	const (
		usage = "Usage: [username] [password]"
	)

	users := map[string]string{
		"jack":  "1888",
		"inanc": "1879",
	}
	var message string

	if len(args) != 3 {
		message = usage
		fmt.Println(message)
	} else {
		password, found := users[args[1]]
		if found {
			if password == args[2] {
				message = "Access granted to " + args[1]
			} else {
				message = "Invalid password for " + args[1]
			}
		} else {
			message = "Access denied for " + args[1]
		}
		fmt.Printf(message)
	}
	return message
}

func main() {
	Authenticate(os.Args)

}
