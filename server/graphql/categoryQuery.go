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
			if n, ok := p.Source.(needJSON); ok {
				id = int(n.CategoryID)
			}

			c, err := get(int64(id))

			if err != nil {
				return nil, err
			}

			cJSON := categoryJSON{
				ID:   c.ID,
				Name: c.Name,
				Icon: c.Icon,
			}

			return cJSON, nil
		},
	}
}
