package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/wawandco/gontact/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/wawandco/gontact/providers"
)

func TestContactValidationFailed(t *testing.T) {
	params := url.Values{
		"Name":  {"Juan"},
		"Email": {"email@email.com"},
	}

	request, _ := http.NewRequest("POST", "/contact", nil)
	request.PostForm = params
	response := httptest.NewRecorder()

	ContactHandler(response, request)
	assert.Equal(t, response.Code, 422)
}

func TestContactValidationSucceeed(t *testing.T) {
	params := url.Values{}
	params.Set("Name", "Juan")
	params.Add("Email", "email@email.com")
	params.Add("Message", "Hey, Hello from here!")

	request, _ := http.NewRequest("POST", "/contact", nil)
	request.PostForm = params
	response := httptest.NewRecorder()
	ContactHandler(response, request)

	assert.Equal(t, response.Code, 201)
}

func TestCallsCorrectProvider(t *testing.T) {
	providers.InternalProviderCount = 0

	params := url.Values{}
	params.Set("Name", "Juan")
	params.Add("Email", "email@email.com")
	params.Add("Message", "Hey, Hello from here!")

	os.Setenv("GONTACT_PROVIDER", "INTERNAL")
	request, _ := http.NewRequest("POST", "/contact", nil)
	request.PostForm = params

	response := httptest.NewRecorder()
	ContactHandler(response, request)

	assert.Equal(t, providers.InternalProviderCount, 1)
	assert.Equal(t, response.Code, 201)
}

func TestCallsCorrectProviderAndReturnsError(t *testing.T) {
	providers.InternalProviderCount = 0

	params := url.Values{}
	params.Set("Name", "Juan")
	params.Add("Email", "email@email.com")
	params.Add("Message", "Hey, Hello from here!")

	os.Setenv("GONTACT_PROVIDER", "ERROR")
	request, _ := http.NewRequest("POST", "/contact", nil)
	request.PostForm = params

	response := httptest.NewRecorder()
	ContactHandler(response, request)

	assert.Equal(t, response.Code, 500)
}
