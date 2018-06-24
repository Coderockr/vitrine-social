package graphql

import "github.com/graphql-go/graphql"

func newSearchQuery() *graphql.Field {
	return &graphql.Field{
		Name:        "SearchQuery",
		Description: "Search active needs on the database",
		Args: graphql.FieldConfigArgument{
			"text":           stringArgument,
			"categories":     intListArgument,
			"organizationId": intArgument,
			"page":           intArgument,
		},
		Type:    graphql.NewList(needType),
		Resolve: graphql.DefaultResolveFn,
	}
}
