package model

import "time"

// User you know it
type User struct {
	Email    string `valid:"email,required" db:"email"`
	Password string `valid:"required" db:"password"`
	ID       int64  `valid:"-"`
}

// Image vinculada a uma necessidade
type image struct {
	ID   int64  `valid:"required" db:"id"`
	Name string `db:"name"`
	URL  string `valid:"required" db:"url"`
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
	Needs   []Need
	Images  []OrganizationImage
}

// OrganizationImage de uma organização
type OrganizationImage struct {
	image
	OrganizationID int64 `valid:"required" db:"organization_id"`
}

// Need uma necessidade da organização
type Need struct {
	ID               int64      `valid:"required" db:"id"`
	Title            string     `valid:"required" db:"title"`
	Description      string     `valid:"required" db:"description"`
	RequiredQuantity int        `db:"required_qtd"`
	ReachedQuantity  int        `db:"reached_qtd"`
	Unity            string     `valid:"required" db:"unity"`
	DueDate          *time.Time `db:"due_date"`
	Status           string     `valid:"required" db:"status"`
	CategoryID       int64      `valid:"required" db:"category_id"`
	OrganizationID   int64      `valid:"required" db:"organization_id"`
	Category         Category
	Images           []NeedImage
}

// NeedImage de uma necessidade
type NeedImage struct {
	image
	NeedID int64 `valid:"required" db:"need_id"`
}

// NeedResponse resposta a necessidade cadastrada da organização
type NeedResponse struct {
	ID      int64      `valid:"required" db:"id"`
	Date    *time.Time `db:"date"`
	Email   string     `valid:"required" db:"email"`
	Name    string     `valid:"required" db:"name"`
	Phone   string     `valid:"required" db:"phone"`
	Address string     `db:"address"`
	Message string     `db:"message"`
	NeedID  int64      `valid:"required" db:"need_id"`
}

// Category de uma necessidade
type Category struct {
	ID   int64  `valid:"required" db:"id"`
	Name string `valid:"required" db:"name"`
	Icon string `valid:"required" db:"icon"`
}
