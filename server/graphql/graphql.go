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
		ChangePassword(o model.Organization, cPassword, nPassword string) (model.Organization, error)
	}

	catRepo interface {
		Get(int64) (*model.Category, error)
		GetAll() ([]model.Category, error)
		GetNeedsCount(*model.Category) (int64, error)
	}

	tokenManager interface {
		CreateToken(model.User) (string, error)
		ValidateToken(token string) (int64, error)
	}
)

func getInArguments(args map[string]interface{}, names ...string) interface{} {
	if len(names) == 0 {
		panic("getInArguments needs at last one name")
	}

	v := args[names[0]]
	if len(names) == 1 {
		return v
	}

	return getInArguments(v.(map[string]interface{}), names[1:]...)
}

// NewHandler returns a handler for the GraphQL implementation of the API
func NewHandler(nR needRepo, oR orgRepo, cR catRepo, tm tokenManager) http.Handler {

	rootQuery := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"search":        newSearchQuery(),
			"need":          newNeedQuery(nR.Get, oR.Get),
			"organization":  newOrganizationQuery(oR.Get),
			"category":      newCategoryQuery(cR.Get, cR.GetNeedsCount),
			"viewer":        newViewerQuery(tm.ValidateToken, oR.Get),
			"allCategories": newAllCategoriesQuery(cR.GetAll),
		},
	}

	rootMutation := graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"login": newLoginMutation(oR.GetUserByEmail, tm.CreateToken, oR.Get),
			"viewer": newViewerMutation(
				tm.ValidateToken,
				oR.Get,
				graphql.Fields{
					"updatePassword": newUpdatePasswordMutation(oR.ChangePassword),
				},
			),
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