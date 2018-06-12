package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gobuffalo/pop/nulls"
)

type jsonTime struct {
	time.Time
}

func (t jsonTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format("\"2006-01-02\"")), nil
}

func (t *jsonTime) UnmarshalJSON(b []byte) (err error) {
	t.Time, err = time.Parse("\"2006-01-02\"", string(b))
	return
}

type baseOrganizationJSON struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Logo  string `json:"logo"`
	Slug  string `json:"slug"`
	Phone string `json:"phone"`
}

type organizationJSON struct {
	baseOrganizationJSON
	About   string      `json:"about"`
	Video   string      `json:"video"`
	Email   string      `json:"email"`
	Address addressJSON `json:"address"`
	Needs   []needJSON  `json:"needs"`
	Images  []imageJSON `json:"images"`
}

type categoryJSON struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type categoryWithCountJSON struct {
	categoryJSON
	NeedsCount int64 `json:"needs_count"`
}

type imageJSON struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type needJSON struct {
	ID               int64                `json:"id"`
	Category         categoryJSON         `json:"category"`
	Organization     baseOrganizationJSON `json:"organization"`
	Images           []imageJSON          `json:"images"`
	Title            string               `json:"title"`
	Description      string               `json:"description"`
	RequiredQuantity int                  `json:"requiredQuantity"`
	ReachedQuantity  int                  `json:"reachedQuantity"`
	Unit             string               `json:"unit"`
	DueDate          *jsonTime            `json:"dueDate"`
	Status           string               `json:"status"`
	CreatedAt        time.Time            `json:"createdAt"`
	UpdatedAt        *time.Time           `json:"updatedAt"`
}

type addressJSON struct {
	Street       string       `json:"street"`
	Number       string       `json:"number"`
	Complement   nulls.String `json:"complement"`
	Neighborhood string       `json:"neighborhood"`
	City         string       `json:"city"`
	State        string       `json:"state"`
	Zipcode      string       `json:"zipcode"`
}

type paginationJSON struct {
	TotalResults int `json:"totalResults"`
	TotalPages   int `json:"totalPages"`
	CurrentPage  int `json:"currentPage"`
}

// SearchResult formato do resultado da busca
type searchResultJSON struct {
	Pagination paginationJSON `json:"pagination"`
	Results    []needJSON     `json:"results"`
}

func requestToJSONObject(req *http.Request, jsonDoc interface{}) error {
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)
	return decoder.Decode(jsonDoc)
}
