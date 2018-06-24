package graphql

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// NewHandler returns a handler for the GraphQL implementation of the API
func NewHandler() http.Handler {

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: graphql.Fields{},
	}

	rootMutation := graphql.ObjectConfig{
		Name:   "RootMutation",
		Fields: graphql.Fields{},
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	})

	if err != nil {
		panic(err)
	}

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}
