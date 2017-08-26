package auth

import "errors"

var (
	ErrEmailDuplication = errors.New("The email is already in the store")
	ErrUserNotFound     = errors.New("User not found")
	ErrWrongPassword    = errors.New("email or password is incorrent")
)

type UserRepository interface {
	Login(email, pass string) (int64, error)
}
