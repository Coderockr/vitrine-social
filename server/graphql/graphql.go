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

	tokenManager interface {
		CreateToken(u model.User, ps *[]string) (string, error)
		ValidateToken(token string) (*model.Token, error)
	}

	catRepo interface {
		Get(int64) (model.Category, error)
	}
)

// NewHandler returns a handler for the GraphQL implementation of the API
func NewHandler(
	nR needRepo,
	oR orgRepo,
	tm tokenManager,
	cR catRepo,
) http.Handler {

	oQuery := newOrganizationQuery(oR.Get)
	cQuery := newCategoryQuery(cR.Get)

	rootQuery := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"need":         newNeedQuery(nR.Get, oR.Get),
			"organization": oQuery,
			"category":     cQuery,
			"viewer":       newViewerQuery(tm.ValidateToken, oR.Get),
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
