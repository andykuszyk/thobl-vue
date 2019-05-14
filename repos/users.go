package repos

import (

	"github.com/andykuszyk/thobl/models"
)

type UsersRepo interface {
	GetByUsername(username string) (*models.User, error)
}

type UsersPostgresRepo struct {

}

func (r *UsersPostgresRepo) GetByUsername(username string) (*models.User, error) {
	return nil, nil
}

type UsersRepoMock struct {
	GetByUsernameFunc func(username string) (*models.User, error)
}

func (r *UsersRepoMock) GetByUsername(username string) (*models.User, error) {
	return r.GetByUsernameFunc(username)
}
