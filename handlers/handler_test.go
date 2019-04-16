package handlers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/stretchr/testify/assert"
)

func TestHandler_PostUser_ShouldReturnToken(t *testing.T) {
	sut := BuildHandlers().Router
	req, err := http.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"username":"spam","password":"eggs"}`))
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	sut.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Posting a user should return a 201")
}
