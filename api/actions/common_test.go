package actions

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// Grab the Gin router with registered routes
func routerWithGetRoute(path string, route gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.GET(path, route)
	return router
}

func routerWithPostRoute(path string, route gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.POST(path, route)
	return router
}

func performGET(r http.Handler, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func performPOST(router http.Handler, path string, payload url.Values) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", path, strings.NewReader(payload.Encode()))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
