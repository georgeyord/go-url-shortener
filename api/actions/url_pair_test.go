package actions

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Grab the Gin router with registered routes
func getRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/api/create", CreateUrlPair)
	return router
}

func TestCreateShouldRespondWithTheSameShortLongUrlsInJsonFormat(t *testing.T) {
	const short = "a7sy9d8"
	const long = "https://www.google.com"

	// Expected body
	expected := gin.H{
		"short": short,
		"long":  long,
	}

	router := routerWithPostRoute("/api/create", CreateUrlPair)

	// Perform a POST request with that handler.
	payload := url.Values{}
	payload.Add("short", short)
	payload.Add("long", long)
	w := performPOST(router, "/api/create", payload)

	// Assert we encoded correctly, the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Grab the value & whether or not it exists
	urlPair, exists := response["urlPair"]

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expected["message"], urlPair)
}
