package graphql

import (
	"errors"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type validateTokenFn func(token string) (*model.Token, error)

func getViewerByToken(get getOrgFn) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		t := getToken(p.Context)
		if t == nil {
			return nil, errors.New("no token was provided, you must send a token through Authorization header")
		}

		return get(t.UserID)
	}
}

func newViewerQuery(get getOrgFn) *graphql.Field {
	return &graphql.Field{
		Name:        "ViewerQuery",
		Description: "Authorized Organization",
		Type:        organizationType,
		Resolve:     getViewerByToken(get),
	}
}
