package handlers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/stretchr/testify/assert"
	"github.com/andykuszyk/thobl/repos"
	"github.com/andykuszyk/thobl/models"
)

func TestHandler_PostUser_ShouldReturnToken(t *testing.T) {
	sut := BuildHandlers(&repos.UsersRepoMock{}).Router
	req, err := http.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"username":"spam","password":"eggs"}`))
	assert.NoError(t, err)
	w := httptest.NewRecorder()

	sut.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Posting a user should return a 201")
}

func TestHandler_PostUsersAuthenticate_ShouldReturn404(t *testing.T) {
	sut := BuildHandlers(&repos.UsersRepoMock{
		GetByUsernameFunc: func (s string) (*models.User, error) {
			return nil, nil
		},
	}).Router
	req, err := http.NewRequest(http.MethodPost, "/users/authenticate", strings.NewReader(`{"username":"spam","password":"eggs"}`))
	assert.NoError(t, err)
	w := httptest.NewRecorder()

	sut.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Trying to authenticate when the user does not exist should return a 404")
}

