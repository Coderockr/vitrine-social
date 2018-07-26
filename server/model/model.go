package model

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"

	"github.com/gobuffalo/pop/nulls"
)

const (
	// PasswordResetPermission allows the user to change its password without sending the current one
	PasswordResetPermission = "password:reset"
)

// Token represents a parsed and validated JWT token
type Token struct {
	UserID      int64
	Permissions map[string]bool
	Token       string
}

// User you know it
type User struct {
	Email    string `valid:"email,required" db:"email"`
	Password string `valid:"required" db:"password"`
	ID       int64  `valid:"-"`
}

// Image vinculada a uma necessidade
type Image struct {
	ID   int64  `valid:"required" db:"id"`
	Name string `db:"name"`
	URL  string `valid:"required" db:"url"`
}

//Organization dados dos usuários que podem logar no sistema
type Organization struct {
	User
	Name        string `valid:"required" db:"name"`
	Logo        *OrganizationImage
	LogoImageID nulls.Int `valid:"optional" db:"logo_image_id"`
	Phone       string    `valid:"required" db:"phone"`
	About       string    `db:"about"`
	Video       string    `valid:"required" db:"video"`
	Slug        string    `valid:"required" db:"slug"`
	Address     Address
	Needs       []Need
	Images      []OrganizationImage
	CreatedAt   *time.Time   `db:"created_at"`
	Website     nulls.String `valid:"optional" db:"website"`
}

// OrganizationImage de uma organização
type OrganizationImage struct {
	Image
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
	Unit             string     `valid:"required" db:"unit"`
	DueDate          *time.Time `db:"due_date"`
	Status           needStatus `valid:"required" db:"status"`
	CategoryID       int64      `valid:"required" db:"category_id"`
	OrganizationID   int64      `valid:"required" db:"organization_id"`
	Category         Category
	Organization     Organization
	Images           []NeedImage
	CreatedAt        time.Time  `db:"created_at"`
	UpdatedAt        *time.Time `db:"updated_at"`
}

// NeedImage de uma necessidade
type NeedImage struct {
	Image
	NeedID int64 `valid:"required" db:"need_id"`
}

// NeedResponse resposta a necessidade cadastrada da organização
type NeedResponse struct {
	ID        int64     `valid:"required" db:"id"`
	Email     string    `valid:"required" db:"email"`
	Name      string    `valid:"required" db:"name"`
	Phone     string    `valid:"required" db:"phone"`
	Address   string    `db:"address"`
	Message   string    `db:"message"`
	NeedID    int64     `valid:"required" db:"need_id"`
	CreatedAt time.Time `valid:"required" db:"created_at"`
}

// Category de uma necessidade
type Category struct {
	ID         int64  `valid:"required" db:"id"`
	Name       string `valid:"required" db:"name"`
	Slug       string `valid:"required" db:"slug"`
	NeedsCount int64  `db:"count_need"`
}

// Address de uma organização
type Address struct {
	Street       string       `valid:"required" db:"street"`
	Number       string       `valid:"required" db:"number"`
	Complement   nulls.String `db:"complement"`
	Neighborhood string       `valid:"required" db:"neighborhood"`
	City         string       `valid:"required" db:"city"`
	State        string       `valid:"required" db:"state"`
	Zipcode      string       `valid:"required" db:"zipcode"`
}

// SearchNeed estrutura de busca de necessidade
type SearchNeed struct {
	Need
	OrganizationName  string       `db:"organization_name"`
	OrganizationLogo  nulls.String `db:"organization_logo"`
	OrganizationSlug  string       `db:"organization_slug"`
	OrganizationPhone string       `db:"organization_phone"`
	CategoryName      string       `db:"category_name"`
	CategorySlug      string       `db:"category_slug"`
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
