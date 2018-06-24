package graphql

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type validateTokenFn func(token string) (*model.Token, error)

func getViewerByToken(validate validateTokenFn, get getOrgFn) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		t, err := validate(p.Args["token"].(string))
		if err != nil {
			return nil, err
		}

		return get(t.UserID)
	}
}

func newViewerQuery(validate validateTokenFn, get getOrgFn) *graphql.Field {
	return &graphql.Field{
		Name:        "ViewerQuery",
		Description: "Authorized Organization",
		Type:        organizationType,
		Args: graphql.FieldConfigArgument{
			"token": nonNullStringArgument,
		},
		Resolve: getViewerByToken(validate, get),
	}
}
