package actions

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Grab the Gin router with registered routes
func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", GetHelloWorld())
	return router
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHelloWorldCallWithNameShouldRespondWithTheNameTransformed(t *testing.T) {
	// Expected body
	const expected = "Hello Foo!"

	router := getRouter()

	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/?name=foo")

	// Assert we encoded correctly, the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Grab the value & whether or not it exists
	message, exists := response["message"]

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expected, message)
}
