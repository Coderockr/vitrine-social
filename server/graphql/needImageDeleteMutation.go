package graphql

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type needImageDeleteFn func(t *model.Token, needID, needImageID int64) error

var needImageDeleteInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "NeedImageDeleteInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"needId":      nonNullIntInput,
		"needImageId": nonNullIntInput,
	},
})

var needImageDeletePayload = graphql.NewObject(graphql.ObjectConfig{
	Name: "NeedImageDeletePayload",
	Fields: graphql.Fields{
		"need": &graphql.Field{
			Type: needType,
		},
	},
})

func newNeedImageDeleteMutation(delete needImageDeleteFn, get getNeedFn) *graphql.Field {
	return &graphql.Field{
		Name:        "NeedImageDeleteMutation",
		Description: "Deletes a image from the need",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: needImageDeleteInput,
			},
		},
		Type: needImageDeletePayload,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			t := getToken(p.Context)
			input := p.Args["input"].(map[string]interface{})

			needID := int64(input["needId"].(int))
			err := delete(
				t,
				needID,
				int64(input["needImageId"].(int)),
			)

			if err != nil {
				return nil, err
			}

			n, _ := get(needID)
			return map[string]interface{}{"need": n}, nil
		},
	}
}
