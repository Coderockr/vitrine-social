package graphql

import (
	"context"
	"errors"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/Coderockr/vitrine-social/server/handlers"
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	graphqlmultipart "github.com/lucassabreu/graphql-multipart-middleware"
)

var errTokenOrgNotFound = errors.New("token organization not found")

type (
	needRepo interface {
		Get(int64) (*model.Need, error)
		GetNeedsImages(model.Need) ([]model.NeedImage, error)
		Create(model.Need) (model.Need, error)
		Update(model.Need) (model.Need, error)
	}

	orgRepo interface {
		Get(int64) (*model.Organization, error)
		Update(model.Organization) (model.Organization, error)
		GetUserByEmail(string) (model.User, error)
		ChangePassword(o model.Organization, cPassword, nPassword string) (model.Organization, error)
		ResetPasswordTo(*model.Organization, string) error
	}

	catRepo interface {
		Get(int64) (*model.Category, error)
		GetAll() ([]model.Category, error)
		GetNeedsCount(*model.Category) (int64, error)
	}

	searchRepo interface {
		Search(text string, categoriesID []int, organizationsID int64, status string, orderBy string, order string, page int) (results []model.SearchNeed, count int, err error)
	}

	tokenManager interface {
		CreateToken(u model.User, ps *[]string) (string, error)
	}

	imageStorage interface {
		CreateNeedImage(*model.Token, int64, *multipart.FileHeader) (*model.NeedImage, error)
	}
)

// NewHandler returns a handler for the GraphQL implementation of the API
func NewHandler(
	nR needRepo,
	oR orgRepo,
	tm tokenManager,
	cR catRepo,
	sR searchRepo,
	iS imageStorage,
) http.Handler {

	rootQuery := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"search":        newSearchQuery(sR.Search),
			"need":          newNeedQuery(nR.Get, oR.Get, nR.GetNeedsImages),
			"organization":  newOrganizationQuery(oR.Get, sR.Search),
			"category":      newCategoryQuery(cR.Get, cR.GetNeedsCount),
			"viewer":        newViewerQuery(oR.Get),
			"allCategories": newAllCategoriesQuery(cR.GetAll),
		},
	}

	rootMutation := graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"login": newLoginMutation(oR.GetUserByEmail, tm.CreateToken, oR.Get),
			"viewer": newViewerMutation(
				oR.Get,
				graphql.Fields{
					"updatePassword":     newUpdatePasswordMutation(oR.ChangePassword),
					"organizationUpdate": newOrganizationUpdateMutation(oR.Update),
					"needCreate":         newNeedCreateMutation(nR.Create),
					"needImageCreate":    newNeedImageCreateMutation(iS.CreateNeedImage),
					"needUpdate":         newNeedUpdateMutation(nR.Get, nR.Update),
					"resetPassword":      newResetPasswordMutation(oR.Get, oR.ResetPasswordTo),
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

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   false,
		GraphiQL: false,
	})

	giqlHandler := &graphiqlHandler{
		next: graphqlmultipart.NewHandler(
			&schema,
			60*1024*1024,
			h,
		),
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		giqlHandler.ServeHTTP(
			w,
			r.WithContext(
				context.WithValue(r.Context(), contextTokenKey, handlers.GetModelToken(r)),
			),
		)
	})
}

type contextKey string

const contextTokenKey = contextKey("tokenKey")

func getToken(c context.Context) *model.Token {
	switch t := c.Value(contextTokenKey).(type) {
	case *model.Token:
		return t
	default:
		return nil
	}
}
