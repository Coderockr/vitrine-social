package repo

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/jmoiron/sqlx"
)

type (
	UserRepository struct {
		DB *sqlx.DB
	}
)

func (r *UserRepository) GetUserByEmail(email string) (model.User, error) {
	u := model.User{}
	err := r.DB.Get(&u, `SELECT * FROM users WHERE email = $1`, email)
	return u, err
}
