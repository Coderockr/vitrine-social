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
			if n, ok := p.Source.(*model.Need); ok {
				id = int(n.CategoryID)
			}

			return get(int64(id))
		},
	}
}
