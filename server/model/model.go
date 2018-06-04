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
	Email    string `valid:"email,required" db:"email" json:"email"`
	Password string `valid:"required" db:"password" json:"password"`
	ID       int64  `valid:"-" json:"id"`
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
	Name      string              `valid:"required" db:"name" json:"name"`
	Logo      string              `valid:"url,optional" db:"logo" json:"logo"`
	Phone     string              `valid:"required" db:"phone" json:"phone"`
	About     string              `db:"about" json:"about"`
	Video     string              `valid:"required" db:"video" json:"video"`
	Slug      string              `valid:"required" db:"slug" json:"slug"`
	Address   Address             `json:"address"`
	Needs     []Need              `json:"needs"`
	Images    []OrganizationImage `json:"images"`
	CreatedAt *time.Time          `db:"created_at" json:"created_at"`
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
	ID               int64      `valid:"required" db:"id" json:"id"`
	Title            string     `valid:"required" db:"title" json:"title"`
	Description      string     `valid:"required" db:"description" json:"description"`
	RequiredQuantity int        `db:"required_qtd" json:"requiredQuantity"`
	ReachedQuantity  int        `db:"reached_qtd" json:"reachedQuantity"`
	Unity            string     `valid:"required" db:"unity" json:"unity"`
	DueDate          *time.Time `db:"due_date" json:"dueDate"`
	Status           needStatus `valid:"required" db:"status" json:"status"`
	CategoryID       int64      `valid:"required" db:"category_id" json:"categoryId"`
	OrganizationID   int64      `valid:"required" db:"organization_id" json:"organizationId"`
	Category         Category
	Images           []NeedImage `json:"images"`
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
	ID         int64  `valid:"required" db:"id" json:"id"`
	Name       string `valid:"required" db:"name" json:"name"`
	Icon       string `valid:"required" db:"icon" json:"icon"`
	NeedsCount int64  `db:"count_need" json:"needs_count"`
}

// Address de uma organização
type Address struct {
	Street       string       `valid:"required" db:"street" json:"street"`
	Number       int64        `valid:"required" db:"number" json:"number"`
	Complement   nulls.String `db:"complement" json:"complement"`
	Neighborhood string       `valid:"required" db:"neighborhood" json:"neighbordhood"`
	City         string       `valid:"required" db:"city" json:"city"`
	State        string       `valid:"required" db:"state" json:"state"`
	Zipcode      string       `valid:"required" db:"zipcode" json:"zipcode"`
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

	switch strings.ToUpper(strings.TrimSpace(str)) {
	case string(NeedStatusActive):
		*s = NeedStatusActive
	case string(NeedStatusInactive):
		*s = NeedStatusInactive
	default:
		*s = NeedStatusEmpty
	}
	return nil
}

func (s needStatus) Value() (driver.Value, error) {
	return string(s), nil
}
