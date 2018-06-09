package graphql

import (
	"errors"
	"log"
	"net/http"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var errTokenOrgNotFound = errors.New("token organization not found")

type (
	needRepo interface {
		Get(int64) (*model.Need, error)
		Create(model.Need) (model.Need, error)
		Update(model.Need) (model.Need, error)
	}

	orgRepo interface {
		Get(int64) (*model.Organization, error)
		Update(model.Organization) (model.Organization, error)
		GetUserByEmail(string) (model.User, error)
		ChangePassword(o model.Organization, cPassword, nPassword string) (model.Organization, error)
	}

	catRepo interface {
		Get(int64) (*model.Category, error)
		GetAll() ([]model.Category, error)
		GetNeedsCount(*model.Category) (int64, error)
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
					"updatePassword":     newUpdatePasswordMutation(oR.ChangePassword),
					"organizationUpdate": newOrganizationUpdateMutation(oR.Update),
					"needCreate":         newNeedCreateMutation(nR.Create),
					"needUpdate":         newNeedUpdateMutation(nR.Get, nR.Update),
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
