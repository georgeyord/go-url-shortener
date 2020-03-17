package actions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/georgeyord/go-url-shortener/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateUrlPairShouldRespondWithTheSameShortLongUrlsInJsonFormat(t *testing.T) {
	const long = "https://www.google.com"
	const short = "a7sy9d8"

	router, _ := runTestRouter()
	// Act
	w := performMockedPost(router, "/api/create", fmt.Sprintf(`{"long":"%s","short": "%s"}`, long, short))

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response to a map
	type Response struct {
		Data models.UrlPair `json:"data"`
	}
	var response Response
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, long, response.Data.Long)
	assert.Equal(t, short, response.Data.Short)
	assert.NotEmpty(t, response.Data.ID)
	assert.NotEmpty(t, response.Data.CreatedAt)
	assert.NotEmpty(t, response.Data.UpdatedAt)
}

func TestCreateUrlPairWithTheSameShortUrtlShouldRespondWithAnErrorInJsonFormat(t *testing.T) {
	const short = "a7sy9d8"
	const long = "https://www.google.com"

	router, _ := runTestRouter()
	w1 := performMockedPost(router, "/api/create", fmt.Sprintf(`{"long":"%s","short": "%s"}`, long, short))
	assert.Equal(t, http.StatusOK, w1.Code)

	// Act
	w2 := performMockedPost(router, "/api/create", fmt.Sprintf(`{"long":"%s","short": "%s"}`, long, short))

	// Assert
	assert.Equal(t, http.StatusBadRequest, w2.Code)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal(w2.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.NotEmpty(t, response["error"])
	assert.Contains(t, response["error"], "Short url is already in use")
}

func TestCreateUrlPairShouldRespondWithTheSameLongAndRadmonShortUrlInJsonFormat(t *testing.T) {
	const long = "https://www.google.com"

	router, _ := runTestRouter()
	// Act
	w := performMockedPost(router, "/api/create", fmt.Sprintf(`{"long":"%s"}`, long))

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response to a map
	type Response struct {
		Data models.UrlPair `json:"data"`
	}
	var response Response
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, long, response.Data.Long)
	assert.NotEmpty(t, response.Data.Short)
	assert.NotEmpty(t, response.Data.ID)
	assert.NotEmpty(t, response.Data.CreatedAt)
	assert.NotEmpty(t, response.Data.UpdatedAt)
}

func TestCreateUrlPairWithoutLongUrlShouldRespondWithErrorInJsonFormat(t *testing.T) {
	router, _ := runTestRouter()
	// Act
	w := performMockedPost(router, "/api/create", "{}")

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.NotEmpty(t, response["error"])
	assert.Contains(t, response["error"], "CreateUrlPairInput.Long", "required")
}

func TestListUrlPairsWithoutAnyPairsStortedShouldRespondWithAnEmptyListOfPairsInJsonFormat(t *testing.T) {
	router, _ := runTestRouter()

	// Act
	w := performMockedGet(router, "/api/list")

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response to a map
	type Response struct {
		Data []models.UrlPair `json:"data"`
	}
	var response Response
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Empty(t, response.Data)
	assert.Len(t, response.Data, 0)
}

func TestListUrlPairsShouldRespondWithAListOfPairsInJsonFormat(t *testing.T) {
	router, db := runTestRouter()

	const long1 = "https://www.google.com"
	const short1 = "9d8a7sy"
	urlPair1 := models.NewUrlPair(long1, short1)
	if err := db.Create(&urlPair1).Error; err != nil {
		log.Fatal(err.Error())
	}
	const long2 = "https://www.google.de"
	const short2 = "97syd8a"
	urlPair2 := models.NewUrlPair(long2, short2)
	if err := db.Create(&urlPair2).Error; err != nil {
		log.Fatal(err.Error())
	}

	// Act
	w := performMockedGet(router, "/api/list")

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response to a map
	type Response struct {
		Data []models.UrlPair `json:"data"`
	}
	var response Response
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Len(t, response.Data, 2)
}

func TestRedirectWithInvalidShortUrlShouldRespondWithRedirectHttpCode(t *testing.T) {
	const long = "https://www.google.com"
	const short = "9d8a7sy"

	router, db := runTestRouter()

	urlPair := models.NewUrlPair(long, short)
	if err := db.Create(&urlPair).Error; err != nil {
		log.Fatal(err.Error())
	}

	// Act
	w := performMockedGet(router, fmt.Sprintf("/r/%s", short))

	// Assert
	assert.Equal(t, http.StatusPermanentRedirect, w.Code)

	response := w.Body.String()
	assert.NotEmpty(t, response, "<a href=\"https://www.google.com\">")
}

func TestRedirectWithInvalidShortUrlShouldRespondWithErrorInJsonFormat(t *testing.T) {
	router, _ := runTestRouter()
	// Act
	w := performMockedGet(router, "/r/doesNotExist")

	// Assert we encoded correctly, the request gives a 400
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Convert the JSON response to a map
	var response map[string]string
	fmt.Println(w.Body.String())
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert
	assert.Nil(t, err)
	assert.NotEmpty(t, response["error"])
	assert.Contains(t, response["error"], "record not found")
}
