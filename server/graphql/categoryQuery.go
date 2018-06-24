package graphql

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type (
	getCatFn func(int64) (model.Category, error)
)

func newCategoryQuery(get getCatFn) *graphql.Field {
	return &graphql.Field{
		Name:        "CategoryQuery",
		Description: "Retrieves a Category by its ID",
		Args: graphql.FieldConfigArgument{
			"id": intArgument,
		},
		Type: categoryType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, _ := p.Args["id"].(int)

			c, err := get(int64(id))

			if err != nil {
				return nil, err
			}

			cJSON := categoryJSON{
				ID:   c.ID,
				Name: c.Name,
				Slug: c.Slug,
			}

			return cJSON, nil
		},
	}
}

type categoryJSON struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
