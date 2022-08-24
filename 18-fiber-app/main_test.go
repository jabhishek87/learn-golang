package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
	description := "index route"
	response_body := "Welcome to Fiber App on Port: 4000"
	// Setup the app as it is done in the main function
	app := Setup()
	req, _ := http.NewRequest("GET", "/", nil)
	res, err := app.Test(req)

	// verify that no error occured, that is not expected
	assert.Equalf(t, false, err != nil, description)

	// Verify if the status code is as expected
	assert.Equalf(t, 200, res.StatusCode, description)

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)

	// Reading the response body should work everytime, such that
	// the err variable should be nil
	assert.Nilf(t, err, description)

	// Verify, that the reponse body equals the expected body
	assert.Equalf(t, response_body, string(body), description)
}
