package main

import "testing"

// run test
// go test -v
// go test - cover
// go test -cover -coverprofile=c.out && go tool cover -html=c.out -o coverage.html

// TestAuthenticate is a func
func TestAuthenticateBlank(t *testing.T) {
	var input []string
	message := Authenticate(input)
	if message != "Usage: [username] [password]" {
		t.Errorf("Test fail ")
	}
}

func TestAuthenticateUser(t *testing.T) {
	input := []string{}
	input = append(input, "/", "Abhishek")
	message := Authenticate(input)
	if message != "Usage: [username] [password]" {
		t.Errorf("Test fail ")
	}
}
