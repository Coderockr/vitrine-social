package inmemory

import (
	"errors"

	"github.com/Coderockr/vitrine-social/server/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	Storage map[string]model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		map[string]model.User{},
	}
}

func (r *UserRepository) Login(email, pass string) (int64, error) {
	for _, u := range r.Storage {
		if u.Email != email {
			continue
		}
		err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
		if err == nil {
			return u.ID, nil
		}
		return 0, errors.New("Login ou senha inválido")
	}
	return 0, errors.New("Login ou senha inválido")
}
