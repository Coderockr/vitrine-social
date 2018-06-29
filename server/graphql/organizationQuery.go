package graphql

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type getOrgFn func(int64) (*model.Organization, error)

var searchOrgNeedsInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "SearchOrganizationNeedsInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"text":       stringInput,
		"categories": intListInput,
		"orderBy": &graphql.InputObjectFieldConfig{
			Type:         orderByEnum,
			DefaultValue: "created_at",
		},
		"order": &graphql.InputObjectFieldConfig{
			Type:         orderEnum,
			DefaultValue: "desc",
		},
		"page": &graphql.InputObjectFieldConfig{
			Type:         graphql.Int,
			DefaultValue: 1,
		},
	},
})

func newOrganizationQuery(get getOrgFn, search searchNeedsFn) *graphql.Field {
	organizationType.AddFieldConfig("needs", newSearchNeedField(
		search,
		func(p graphql.ResolveParams) (searchParams, error) {
			o := p.Source.(*model.Organization)

			sp := searchParams{
				OrganizationID: o.ID,
			}

			input, ok := p.Args["input"].(map[string]interface{})
			if !ok {

				return sp, nil
			}

			sp.Text, _ = input["text"].(string)
			sp.Categories, _ = input["categories"].([]int)
			sp.OrderBy, _ = input["orderBy"].(string)
			sp.Order, _ = input["order"].(string)
			sp.Page, _ = input["page"].(int)

			return sp, nil
		},
		graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: searchOrgNeedsInput,
			},
		},
	))

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
