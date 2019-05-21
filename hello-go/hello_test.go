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

		assert.Equal(t, response.Code, http.StatusOK)
		assert.Equal(t, response.Body.String(), "Hello, Fred")
	})
}