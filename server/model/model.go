package model

// User you know it
type User struct {
	Email    string `valid:"email,required" db:"email"`
	Password string `valid:"required" db:"password"`
	ID       int64  `valid:"-"`
}

//Organization dados dos usuários que podem logar no sistema
type Organization struct {
	User
	Name    string `valid:"required" db:"name"`
	Logo    string `valid:"url,optional" db:"logo"`
	Address string `valid:"required" db:"address"`
	Phone   string `valid:"required" db:"phone"`
	Resume  string `db:"resume"`
	Video   string `valid:"required" db:"video"`
	Slug    string `valid:"required" db:"slug"`
}
