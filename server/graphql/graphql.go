package graphql

import (
	"net/http"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type (
	needRepo interface {
		Get(int64) (*model.Need, error)
	}

	orgRepo interface {
		Get(int64) (*model.Organization, error)
		GetUserByEmail(string) (model.User, error)
	}

	catRepo interface {
		Get(int64) (model.Category, error)
	}

	tokenManager interface {
		CreateToken(model.User) (string, error)
	}
)

// NewHandler returns a handler for the GraphQL implementation of the API
func NewHandler(nR needRepo, oR orgRepo, cR catRepo, tm tokenManager) http.Handler {

	oQuery := newOrganizationQuery(oR.Get)
	cQuery := newCategoryQuery(cR.Get)

	rootQuery := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"search":       newSearchQuery(),
			"need":         newNeedQuery(nR.Get, oR.Get),
			"organization": oQuery,
			"category":     cQuery,
		},
	}

	rootMutation := graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"login": newLoginMutation(oR.GetUserByEmail, tm.CreateToken),
		},
	}

	needType.AddFieldConfig("organization", oQuery)
	needType.AddFieldConfig("category", cQuery)

	loginType.AddFieldConfig("organization", oQuery)

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
