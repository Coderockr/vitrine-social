package graphql

import (
	"net/http"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type (
	orgRepo interface {
		Get(int64) (*model.Organization, error)
		GetUserByEmail(string) (model.User, error)
	}

	tokenManager interface {
		CreateToken(u model.User, ps *[]string) (string, error)
		ValidateToken(token string) (*model.Token, error)
	}
)

// NewHandler returns a handler for the GraphQL implementation of the API
func NewHandler(
	oR orgRepo,
	tm tokenManager,
) http.Handler {

	oQuery := newOrganizationQuery(oR.Get)

	rootQuery := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"organization": oQuery,
			"viewer":       newViewerQuery(tm.ValidateToken, oR.Get),
		},
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
