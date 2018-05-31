package graphql

import (
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
			return get(int64(id))
		},
	}
}
