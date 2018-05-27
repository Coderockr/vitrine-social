package graphql

import (
	"time"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type (
	getNeedFn func(int64) (*model.Need, error)
)

func newNeedQuery(get getNeedFn, getOrg getOrgFn) *graphql.Field {
	return &graphql.Field{
		Name:        "NeedQuery",
		Description: "Retrieves a Need by its Id",
		Args: graphql.FieldConfigArgument{
			"id": nonNullIntArgument,
		},
		Type: needType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(int)

			n, err := get(int64(id))
			if err != nil {
				return nil, err
			}

			var dueDate *jsonTime
			if n.DueDate != nil {
				dueDate = &jsonTime{*n.DueDate}
			}

			needJSON := needJSON{
				ID:               n.ID,
				Title:            n.Title,
				Description:      n.Description,
				RequiredQuantity: n.RequiredQuantity,
				ReachedQuantity:  n.ReachedQuantity,
				Unity:            n.Unity,
				DueDate:          dueDate,
				CategoryID:       n.Category.ID,
				OrganizationID:   n.OrganizationID,
				// Images: needImagesToImageJSON(n.Images),
				Status: string(n.Status),
			}

			return needJSON, err
		},
	}
}

type needJSON struct {
	ID               int64       `json:"id"`
	CategoryID       int64       `json:"categoryId"`
	OrganizationID   int64       `json:"organizationId"`
	Images           []imageJSON `json:"images"`
	Title            string      `json:"title"`
	Description      string      `json:"description"`
	RequiredQuantity int         `json:"requiredQuantity"`
	ReachedQuantity  int         `json:"reachedQuantity"`
	Unity            string      `json:"unity"`
	DueDate          *jsonTime   `json:"dueDate"`
	Status           string      `json:"status"`
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

type baseOrganizationJSON struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
	Slug string `json:"slug"`
}

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
