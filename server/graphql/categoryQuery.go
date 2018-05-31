package graphql

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type (
	getCatFn func(int64) (*model.Category, error)
)

func newCategoryQuery(get getCatFn) *graphql.Field {
	f := newCategoryField(func(p graphql.ResolveParams) (*model.Category, error) {
		if id, ok := p.Args["id"].(int); ok {
			return get(int64(id))
		}
		return nil, nil
	})

	f.Name = "CategoryQuery"
	f.Description = "Retrieves a Category by its ID"
	f.Args = graphql.FieldConfigArgument{
		"id": nonNullIntArgument,
	}
	return f
}

type getCatByResolveParams func(graphql.ResolveParams) (*model.Category, error)

func newCategoryField(get getCatByResolveParams) *graphql.Field {
	return &graphql.Field{
		Type: categoryType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return get(p)
		},
	}
}
