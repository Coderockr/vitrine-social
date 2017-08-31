package auth

type UserRepository interface {
	Login(email, pass string) (int64, error)
}
