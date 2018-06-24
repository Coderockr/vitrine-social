package graphql

import (
	"log"
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
		Get(int64) (*model.Category, error)
		GetAll() ([]model.Category, error)
	}

	tokenManager interface {
		CreateToken(u model.User, ps *[]string) (string, error)
		ValidateToken(token string) (*model.Token, error)
	}
)

// NewHandler returns a handler for the GraphQL implementation of the API
func NewHandler(
	nR needRepo,
	oR orgRepo,
	tm tokenManager,
	cR catRepo,
) http.Handler {

	rootQuery := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"search":        newSearchQuery(),
			"need":          newNeedQuery(nR.Get, oR.Get),
			"organization":  newOrganizationQuery(oR.Get),
			"category":      newCategoryQuery(cR.Get),
			"viewer":        newViewerQuery(tm.ValidateToken, oR.Get),
			"allCategories": newAllCategoriesQuery(cR.GetAll),
		},
	}

	rootMutation := graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"login": newLoginMutation(oR.GetUserByEmail, tm.CreateToken, oR.Get),
		},
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	})

	if err != nil {
		log.Fatal(err)
	}

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}
