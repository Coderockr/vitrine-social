package graphql

import (
	"errors"
	"mime/multipart"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type needImageCreateFn func(*model.Token, int64, *multipart.FileHeader) (*model.NeedImage, error)

var needImageCreateInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "NeedImageCreateInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"needId": nonNullIntInput,
		"file":   nonNullUploadInput,
	},
})

var needImageCreatePayload = graphql.NewObject(graphql.ObjectConfig{
	Name: "NeedImageCreatePayload",
	Fields: graphql.Fields{
		"needImage": &graphql.Field{
			Type: needImageType,
		},
	},
})

func newNeedImageCreateMutation(create needImageCreateFn) *graphql.Field {
	return &graphql.Field{
		Name:        "NeedImageCreateMutation",
		Description: "Creates a image for the need",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: needImageCreateInput,
			},
		},
		Type: needImageCreatePayload,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if _, ok := p.Source.(*model.Organization); !ok {
				return nil, errors.New("no organization found")
			}

			t := getToken(p.Context)

			input := p.Args["input"].(map[string]interface{})

			i, err := create(
				t,
				int64(input["needId"].(int)),
				input["file"].(*multipart.FileHeader),
			)

			if err != nil {
				return nil, err
			}

			return map[string]interface{}{"needImage": *i}, nil
		},
	}
}
