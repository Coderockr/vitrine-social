package graphql

import (
	"time"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type needCreateFn func(model.Need) (model.Need, error)

var needCreateInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "NeedCreateInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"title":       nonNullStringInput,
		"description": nonNullStringInput,
		"requiredQuantity": &graphql.InputObjectFieldConfig{
			Type:         graphql.Int,
			DefaultValue: 0,
		},
		"reachedQuantity": &graphql.InputObjectFieldConfig{
			Type:         graphql.Int,
			DefaultValue: 0,
		},
		"unit":       nonNullStringInput,
		"dueDate":    dateInput,
		"categoryId": nonNullIntInput,
	},
})

var needCreatePayload = graphql.NewObject(graphql.ObjectConfig{
	Name: "NeedCreatePayload",
	Fields: graphql.Fields{
		"need": &graphql.Field{
			Type: needType,
		},
	},
})

func newNeedCreateMutation(create needCreateFn) *graphql.Field {
	return &graphql.Field{
		Name:        "NeedCreateMutation",
		Description: "Creates a need for the token Organization",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: needCreateInput,
			},
		},
		Type: needCreatePayload,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			o := p.Source.(*model.Organization)
			input := p.Args["input"].(map[string]interface{})

			n := model.Need{
				OrganizationID:   o.ID,
				CategoryID:       int64(input["categoryId"].(int)),
				Title:            input["title"].(string),
				Description:      input["description"].(string),
				Unit:             input["unit"].(string),
				RequiredQuantity: input["requiredQuantity"].(int),
				ReachedQuantity:  input["reachedQuantity"].(int),
			}

			if dueDate, ok := input["dueDate"].(time.Time); ok {
				n.DueDate = &dueDate
			}

			n, err := create(n)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{"need": &n}, nil
		},
	}
}
