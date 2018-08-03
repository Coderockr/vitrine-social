package graphql

import "github.com/graphql-go/graphql"

func newViewerMutation(get getOrgFn, fields graphql.Fields) *graphql.Field {
	viewerMutationType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "ViewerMutations",
		Fields: fields,
	})

	return &graphql.Field{
		Name:        "Viewer mutations",
		Description: "Mutations that need authentication",
		Type:        viewerMutationType,
		Resolve:     getViewerByToken(get),
	}
}
