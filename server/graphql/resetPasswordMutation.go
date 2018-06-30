package graphql

import (
	"errors"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type resetPasswordToFn func(*model.Organization, string) error

var resetPasswordPayload = graphql.NewObject(graphql.ObjectConfig{
	Name: "ResetPasswordPayload",
	Fields: graphql.Fields{
		"organization": &graphql.Field{
			Type: organizationType,
		},
	},
})

func newResetPasswordMutation(validate validateTokenFn, get getOrgFn, reset resetPasswordToFn) *graphql.Field {
	return &graphql.Field{
		Name:        "ResetPasswordMutation",
		Description: "Reset the Organizations password to a new password",
		Args: graphql.FieldConfigArgument{
			"token":       nonNullStringArgument,
			"newPassword": nonNullStringArgument,
		},
		Type: resetPasswordPayload,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			t, err := validate(p.Args["token"].(string))
			if err != nil {
				return nil, err
			}

			if !t.Permissions[model.PasswordResetPermission] {
				return nil, errors.New("token does not have permission to reset the password")
			}

			o, _ := get(t.UserID)
			if err := reset(o, p.Args["newPassword"].(string)); err != nil {
				return nil, err
			}

			return map[string]interface{}{"organization": o}, nil
		},
	}
}
