package actions

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/georgeyord/go-url-shortener/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func runTestRouter() (*gin.Engine, *gorm.DB) {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	db := initTestDb()
	SetupMiddlewares(router, db)
	MapRoutes(router)
	return router, db
}

// Grab the Gin router with registered routes
func initTestDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err.Error())
	}
	models.SetupModels(db)
	return db
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
		panic(err.Error())
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
