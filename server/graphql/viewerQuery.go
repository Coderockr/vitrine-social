package graphql

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type validateToken func(token string) (*model.Token, error)

func newViewerQuery(validate validateToken, get getOrgFn) *graphql.Field {
	return &graphql.Field{
		Name:        "ViewerQuery",
		Description: "Authorized Organization",
		Type:        organizationType,
		Args: graphql.FieldConfigArgument{
			"token": nonNullStringArgument,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			t, err := validate(p.Args["token"].(string))
			if err != nil {
				return nil, err
			}

			o, err := get(t.UserID)
			if err != nil {
				return nil, err
			}

			return orgToJSON(o), nil
		},
	}
}
