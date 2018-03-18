package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type jsonTime struct {
	time.Time
}

func (t jsonTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Format(time.RFC3339))), nil
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
