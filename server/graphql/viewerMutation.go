package graphql

import "github.com/graphql-go/graphql"

func newViewerMutation(validate validateTokenFn, get getOrgFn, fields graphql.Fields) *graphql.Field {
	viewerMutationType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "ViewerMutations",
		Fields: fields,
	})

	return &graphql.Field{
		Name:        "Viewer mutations",
		Description: "Mutations that need authentication",
		Args: graphql.FieldConfigArgument{
			"token": nonNullStringArgument,
		},
		Type:    viewerMutationType,
		Resolve: getViewerByToken(validate, get),
	}
}
