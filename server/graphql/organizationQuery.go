package graphql

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type getOrgFn func(int64) (*model.Organization, error)

func newOrganizationQuery(get getOrgFn) *graphql.Field {
	return &graphql.Field{
		Name:        "OrganizationQuery",
		Description: "Retrieves a Organization by its Id",
		Args: graphql.FieldConfigArgument{
			"id": intArgument,
		},
		Type: organizationType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, _ := p.Args["id"].(int)
			if n, ok := p.Source.(*model.Need); ok {
				id = int(n.OrganizationID)
			}
			if l, ok := p.Source.(loginJSON); ok {
				id = int(l.OrganizationID)
			}

			return get(int64(id))
		},
	}
}
