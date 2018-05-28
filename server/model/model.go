package model

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"

	"github.com/gobuffalo/pop/nulls"
)

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
	Name   string `valid:"required" db:"name"`
	Logo   string `valid:"url,optional" db:"logo"`
	Phone  string `valid:"required" db:"phone"`
	Resume string `db:"resume"`
	Video  string `valid:"required" db:"video"`
	Slug   string `valid:"required" db:"slug"`
	Address
	Needs     []Need
	Images    []OrganizationImage
	CreatedAt *time.Time `db:"created_at"`
}

// OrganizationImage de uma organização
type OrganizationImage struct {
	image
	OrganizationID int64 `valid:"required" db:"organization_id"`
}

type needStatus string

var (
	// NeedStatusActive a active need
	NeedStatusActive = needStatus("ACTIVE")
	// NeedStatusInactive a inactive need
	NeedStatusInactive = needStatus("INACTIVE")
	// NeedStatusEmpty was not informed
	NeedStatusEmpty = needStatus("")
)

// Need uma necessidade da organização
type Need struct {
	ID               int64      `valid:"required" db:"id"`
	Title            string     `valid:"required" db:"title"`
	Description      string     `valid:"required" db:"description"`
	RequiredQuantity int        `db:"required_qtd"`
	ReachedQuantity  int        `db:"reached_qtd"`
	Unity            string     `valid:"required" db:"unity"`
	DueDate          *time.Time `db:"due_date"`
	Status           needStatus `valid:"required" db:"status"`
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
	Date    *time.Time `valid:"required" db:"date"`
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

// Address de uma organização
type Address struct {
	Street     string       `valid:"required" db:"street"`
	Number     int64        `valid:"required" db:"number"`
	Complement nulls.String `db:"complement"`
	Suburb     string       `valid:"required" db:"suburb"`
	City       string       `valid:"required" db:"city"`
	State      string       `valid:"required" db:"state"`
	Zipcode    string       `valid:"required" db:"zipcode"`
}

func (s *needStatus) Scan(src interface{}) error {
	var str string

	switch src.(type) {
	case string:
		str = src.(string)
	case []byte:
		str = string(src.([]byte))
	default:
		return errors.New("Incompatible type for needStatus")
	}

	switch strings.ToUpper(str) {
	case string(NeedStatusActive):
		s = &NeedStatusActive
	case string(NeedStatusInactive):
		s = &NeedStatusInactive
	default:
		s = &NeedStatusEmpty
	}

	return nil
}

func (s needStatus) Value() (driver.Value, error) {
	return string(s), nil
}
