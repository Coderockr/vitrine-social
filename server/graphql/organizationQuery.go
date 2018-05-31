package graphql

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type getOrgFn func(int64) (*model.Organization, error)

func newOrganizationQuery(get getOrgFn) *graphql.Field {
	f := newOrganizationField(func(p graphql.ResolveParams) (*model.Organization, error) {
		if id, ok := p.Args["id"].(int); ok {
			return get(int64(id))
		}
		return nil, nil
	})

	f.Args = graphql.FieldConfigArgument{
		"id": nonNullIntArgument,
	}
	f.Name = "OrganizationQuery"
	f.Description = "Retrieves a Organization by its Id"

	return f
}

type getOrgByResolveParams func(graphql.ResolveParams) (*model.Organization, error)

func newOrganizationField(get getOrgByResolveParams) *graphql.Field {
	return &graphql.Field{
		Type: organizationType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return get(p)
		},
	}
}
