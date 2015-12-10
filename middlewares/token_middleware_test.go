package middlewares

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/wawandco/gontact/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestTokenBlank(t *testing.T) {
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/contact", nil)
	TokenMiddleware(rw, req, afterMiddlewareFunc)
	assert.Equal(t, myvar, 0)
}

func TestTokenInvalid(t *testing.T) {
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/contact", nil)
	os.Setenv("GONTACT_TOKEN", "AAB")
	req.Header["X-GONTACT-TOKEN"] = []string{"AAA"}
	TokenMiddleware(rw, req, afterMiddlewareFunc)
	assert.Equal(t, myvar, 0)
}

func TestTokenValid(t *testing.T) {
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/contact", nil)
	os.Setenv("GONTACT_TOKEN", "AAB")
	req.Header["X-GONTACT-TOKEN"] = []string{"AAB"}
	TokenMiddleware(rw, req, afterMiddlewareFunc)
	assert.Equal(t, myvar, 100)
}

func TestTokenEmptyEnvEmpty(t *testing.T) {
	myvar = 0
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/contact", nil)
	os.Unsetenv("GONTACT_TOKEN")

	TokenMiddleware(rw, req, afterMiddlewareFunc)
	assert.Equal(t, myvar, 0)
}

var myvar = 0

func afterMiddlewareFunc(rw http.ResponseWriter, r *http.Request) {
	myvar = 100
}
