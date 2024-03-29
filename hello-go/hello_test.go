package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGETHello(t *testing.T) {
	server := &HelloServer{}

	t.Run("greet Fred", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/hello/%s", "Fred"), nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Hello, Fred\n", response.Body.String())
	})

	t.Run("greet Wilma", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/hello/%s", "Wilma"), nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Hello, Wilma\n", response.Body.String())
	})

	t.Run("greet with no path", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Hello\n", response.Body.String())
	})

	t.Run("greet with long path", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/hello/hello/is/anybody/in/there", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Hello, hello/is/anybody/in/there\n", response.Body.String())
	})

	t.Run("call non greeting path", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/is/anybody/in/there", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}