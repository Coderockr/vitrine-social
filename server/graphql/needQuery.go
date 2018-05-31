package graphql

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type (
	getNeedFn func(int64) (*model.Need, error)
)

func newNeedQuery(get getNeedFn, getOrg getOrgFn) *graphql.Field {

	needType.AddFieldConfig(
		"organization",
		newOrganizationField(func(p graphql.ResolveParams) (*model.Organization, error) {
			if n, ok := p.Source.(*model.Need); ok && n != nil {
				return getOrg(n.OrganizationID)
			}

			return nil, nil
		}),
	)

	needType.AddFieldConfig(
		"category",
		newCategoryField(func(p graphql.ResolveParams) (*model.Category, error) {
			if n, ok := p.Source.(*model.Need); ok && n != nil {
				return &n.Category, nil
			}
			return nil, nil
		}),
	)

	return &graphql.Field{
		Name:        "NeedQuery",
		Description: "Retrieves a Need by its Id",
		Args: graphql.FieldConfigArgument{
			"id": nonNullIntArgument,
		},
		Type: needType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if id, ok := p.Args["id"].(int); ok {
				return get(int64(id))
			}
			return nil, nil
		},
	}
}
