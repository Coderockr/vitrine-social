package graphql

import "github.com/graphql-go/graphql"

type validateTokenFn func(token string) (userId int64, err error)

func getViewerByToken(validate validateTokenFn, get getOrgFn) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, err := validate(p.Args["token"].(string))
		if err != nil {
			return nil, err
		}

		return get(id)
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
