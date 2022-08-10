package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// sample only
func TestPingRoute(t *testing.T) {
	router := initRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, Ping.Path, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
