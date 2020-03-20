package actions

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/georgeyord/go-url-shortener/pkg/test/common"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rs/zerolog/log"
)

func runTestRouter() (*gin.Engine, *gorm.DB) {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	db := common.InitTestDb()
	kafkaWriters := common.InitTestKafkaWriters()
	SetupMiddlewares(router, db, kafkaWriters)
	MapRoutes(router)
	return router, db
}

func performMockedGet(r http.Handler, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func performMockedPost(router http.Handler, path string, payload string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(http.MethodPost, path, strings.NewReader(payload))

	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
