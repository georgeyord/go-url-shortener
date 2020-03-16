package actions

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldCallWithNameShouldRespondWithTheNameTransformed(t *testing.T) {
	// Expected body
	const expected = "Hello Foo!"

	router, _ := runTestRouter()

	// Perform a GET request with that handler.
	w := performMockedGet(router, "/?name=foo")

	// Assert we encoded correctly, the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the value & whether or not it exists
	message, exists := response["message"]

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expected, message)
}

func TestHelloWorldCallWithoutNameShouldRespondWithError(t *testing.T) {
	// Expected body
	const expected = "Hello Foo!"

	router, _ := runTestRouter()

	// Perform a GET request with that handler.
	w := performMockedGet(router, "/")

	// Assert we encoded correctly, the request gives a 200
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.NotEmpty(t, response["error"])
	assert.Contains(t, response["error"], "name", "missing")
}
