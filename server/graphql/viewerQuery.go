package graphql

import "github.com/graphql-go/graphql"

type validateToken func(token string) (userId int64, err error)

func newViewerQuery(validate validateToken, get getOrgFn) *graphql.Field {
	return &graphql.Field{
		Name:        "ViewerQuery",
		Description: "Authorized Organization",
		Type:        organizationType,
		Args: graphql.FieldConfigArgument{
			"token": nonNullStringArgument,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, err := validate(p.Args["token"].(string))
			if err != nil {
				return nil, err
			}

			return get(id)
		},
	}
}
