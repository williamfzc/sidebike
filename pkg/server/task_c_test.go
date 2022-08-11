package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// sample only
func TestPingRoute(t *testing.T) {
	engine := gin.Default()
	InitRouter(engine)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, Ping.Path, nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
