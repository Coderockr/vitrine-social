package db

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) Login(email, pass string) (int64, error) {
	var user model.User
	err := r.db.QueryRowx("SELECT o.id, o.email, o.password FROM organizations o WHERE o.email = $1", email).StructScan(&user)
	if err != nil {
		log.Printf("[WARN][VITRINE] %s", err)
		return 0, errors.New("Login ou senha inválido")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil {
		return 0, errors.New("Login ou senha inválido")
	}
	return user.ID, nil
}
