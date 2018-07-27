package graphql

import (
	"errors"
	"log"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/Coderockr/vitrine-social/server/security"
	"github.com/graphql-go/graphql"
)

type (
	getUserByEmail func(email string) (model.User, error)
	createToken    func(u model.User, ps *[]string) (string, error)
)

var loginType = graphql.NewObject(graphql.ObjectConfig{
	Name: "LoginResult",
	Fields: graphql.Fields{
		"token": nonNullStringField,
	},
})

type loginJSON struct {
	Token          string `json:"token"`
	OrganizationID int64  `json:"organizationId"`
}

func newLoginMutation(get getUserByEmail, cT createToken, getOrg getOrgFn) *graphql.Field {

	loginType.AddFieldConfig(
		"organization",
		newOrganizationField(func(p graphql.ResolveParams) (*model.Organization, error) {
			l, _ := p.Source.(loginJSON)
			return getOrg(l.OrganizationID)
		}),
	)

	return &graphql.Field{
		Name:        "LoginMutation",
		Description: "Authenticate the user and returns a token and organization if succeded",
		Type:        loginType,
		Args: graphql.FieldConfigArgument{
			"email":    nonNullStringArgument,
			"password": nonNullStringArgument,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			email := p.Args["email"].(string)
			pass := p.Args["password"].(string)

			user, err := get(email)
			if err != nil {
				log.Printf("[INFO][Login Mutation] %s", err.Error())
				return nil, err
			}
			err = security.CompareHashAndPassword(user.Password, pass)
			if err != nil {
				log.Printf("[INFO][Login Mutation] %s", err.Error())
				return nil, errors.New("password does not match")
			}

			token, err := cT(user, nil)
			if err != nil {
				return nil, err
			}

			result := loginJSON{
				Token:          token,
				OrganizationID: user.ID,
			}
			return result, nil
		},
	}
}
