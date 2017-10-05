package model

import "time"

// User you know it
type User struct {
	Email    string `valid:"email,required" db:"email"`
	Password string `valid:"required" db:"password"`
	ID       int64  `valid:"-"`
}

//Need
type Need struct {
	OrganizationID int64     `db:"organization_id"`
	Title          string    `db:"title"`
	Description    string    `db:"description"`
	RequiredQTD    int64     `db:"required_qtd"`
	ReachedQTD     int64     `db:"reached_qtd"`
	DueDate        time.Time `db:"due_date"`
	Status         string    `db:"status"`
	Unity          string    `db:"unity"`
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
