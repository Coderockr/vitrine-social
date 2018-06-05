package graphql

import (
	"errors"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type changePasswordFn func(o model.Organization, cPassword, nPassword string) (model.Organization, error)

var updatePasswordInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdatePasswordInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"currentPassword": nonNullStringInput,
		"newPassword":     nonNullStringInput,
	},
})

func newUpdatePasswordMutation(cp changePasswordFn) *graphql.Field {
	return &graphql.Field{
		Name:        "UpdatePasswordMutation",
		Description: "Updates the current user password",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: updatePasswordInput,
			},
		},
		Type: organizationType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if o, ok := p.Source.(*model.Organization); ok && o != nil {
				args := p.Args["input"].(map[string]interface{})

				o, err := cp(
					*o,
					args["currentPassword"].(string),
					args["newPassword"].(string),
				)

				if err != nil {
					return nil, err
				}

				return o, nil
			}
			return nil, errors.New("organization not found")
		},
	}
}
