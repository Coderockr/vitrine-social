package handlers

import (
	"encoding/json"
	"net/http"
	"time"
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
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
	Slug string `json:"slug"`
}

type organizationJSON struct {
	baseOrganizationJSON
	Address string      `json:"address"`
	Phone   string      `json:"phone"`
	Resume  string      `json:"resume"`
	Video   string      `json:"video"`
	Email   string      `json:"email"`
	Needs   []needJSON  `json:"needs"`
	Images  []imageJSON `json:"images"`
}

type categoryJSON struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
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
	Unity            string               `json:"unity"`
	DueDate          *jsonTime            `json:"dueDate"`
	Status           string               `json:"status"`
}

func requestToJSONObject(req *http.Request, jsonDoc interface{}) error {
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)
	return decoder.Decode(jsonDoc)
}
