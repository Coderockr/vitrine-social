package graphql

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type orgImageDeleteFn func(t *model.Token, organizationImageID int64) error

var organizationImageDeleteInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "OrganizationImageDeleteInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"organizationImageId": nonNullIntInput,
	},
})

var organizationImageDeletePayload = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrganizationImageDeletePayload",
	Fields: graphql.Fields{
		"organization": &graphql.Field{
			Type: organizationType,
		},
	},
})

func newOrganizationImageDeleteMutation(delete orgImageDeleteFn, get getOrgFn) *graphql.Field {
	return &graphql.Field{
		Name:        "OrganizationImageDeleteMutation",
		Description: "Deletes a image from the current organization",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: organizationImageDeleteInput,
			},
		},
		Type: organizationImageDeletePayload,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			t := getToken(p.Context)
			input := p.Args["input"].(map[string]interface{})

			if err := delete(t, int64(input["organizationImageId"].(int))); err != nil {
				return nil, err
			}

			o, _ := get(t.UserID)
			return map[string]interface{}{"organization": o}, nil
		},
	}
}
