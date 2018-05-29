package inmemory

import (
	"errors"

	"github.com/Coderockr/vitrine-social/server/model"
)

type UserRepository struct {
	Storage map[string]model.User
}

func (r *UserRepository) GetUserByEmail(email string) (model.User, error) {
	for _, u := range r.Storage {
		if u.Email == email {
			return u, nil
		}
	}
	return model.User{}, errors.New("Login ou senha inválido")
}
