package graphql

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type getAllCatsFn func() ([]model.Category, error)

func newAllCategoriesQuery(getAllCat getAllCatsFn) *graphql.Field {
	return &graphql.Field{
		Name:        "AllCategoriesQuery",
		Description: "Return all Categories",
		Type:        graphql.NewList(categoryType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return getAllCat()
		},
	}
}
