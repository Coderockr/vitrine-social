package graphql

import (
	"errors"
	"mime/multipart"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type orgImageCreateFn func(*model.Token, *multipart.FileHeader) (*model.OrganizationImage, error)

var organizationImageCreateInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "OrganizationImageCreateInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"file": nonNullUploadInput,
	},
})

var organizationImageCreatePayload = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrganizationImageCreatePayload",
	Fields: graphql.Fields{
		"organizationImage": &graphql.Field{
			Type: organizationImageType,
		},
	},
})

func newOrganizationImageCreateMutation(create orgImageCreateFn) *graphql.Field {
	return &graphql.Field{
		Name:        "OrganizationImageCreateMutation",
		Description: "Creates a image for the current organization",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: organizationImageCreateInput,
			},
		},
		Type: organizationImageCreatePayload,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if _, ok := p.Source.(*model.Organization); !ok {
				return nil, errors.New("no organization found")
			}

			t := getToken(p.Context)

			input := p.Args["input"].(map[string]interface{})

			i, err := create(t, input["file"].(*multipart.FileHeader))

			if err != nil {
				return nil, err
			}

			return map[string]interface{}{"organizationImage": *i}, nil
		},
	}
}
