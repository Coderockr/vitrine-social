package graphql

import (
	"errors"
	"fmt"
	"time"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gobuffalo/pop/nulls"
	"github.com/graphql-go/graphql"
)

type needUpdateFn func(model.Need) (model.Need, error)

var needPatchInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "NeedPatchInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"title":            stringInput,
		"description":      stringInput,
		"requiredQuantity": intInput,
		"reachedQuantity":  intInput,
		"unit":             stringInput,
		"dueDate":          dateInput,
		"categoryId":       intInput,
		"status":           needStatusInput,
		"withoutDescription": &graphql.InputObjectFieldConfig{
			Type:         graphql.Boolean,
			Description:  "When set to true, will make the need without a description",
			DefaultValue: false,
		},
		"withoutDueDate": &graphql.InputObjectFieldConfig{
			Type:         graphql.Boolean,
			Description:  "When set to true, will make the need without date limit",
			DefaultValue: false,
		},
	},
})

var needUpdateInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "NeedUpdateInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": nonNullIntInput,
		"patch": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(needPatchInput),
		},
	},
})

var needUpdatePayload = graphql.NewObject(graphql.ObjectConfig{
	Name: "NeedUpdatePayload",
	Fields: graphql.Fields{
		"need": &graphql.Field{
			Type: needType,
		},
	},
})

func newNeedUpdateMutation(get getNeedFn, update needUpdateFn) *graphql.Field {
	return &graphql.Field{
		Name:        "NeedUpdateMutation",
		Description: "Updates a need with the patch",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: needUpdateInput,
			},
		},
		Type: needUpdatePayload,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			o, ok := p.Source.(*model.Organization)
			if !ok {
				return nil, errTokenOrgNotFound
			}

			input := p.Args["input"].(map[string]interface{})
			n, err := get(int64(input["id"].(int)))
			if err != nil {
				return nil, err
			}

			if n.OrganizationID != o.ID {
				return nil, fmt.Errorf(
					"organization %d does not own need %d",
					o.ID,
					n.ID,
				)
			}

			patch := input["patch"].(map[string]interface{})

			if title, ok := patch["title"].(string); ok {
				n.Title = title
			}

			withoutDescription := patch["withoutDescription"].(bool)
			if withoutDescription {
				n.Description = nulls.String{Valid: false}
			}

			if description, ok := patch["description"].(string); ok {
				if withoutDescription {
					return nil, errors.New("parameters withoutDescription and description can't be used together")
				}
				n.Description = nulls.NewString(description)
			}

			if unit, ok := patch["unit"].(string); ok {
				n.Unit = unit
			}

			if requiredQuantity, ok := patch["requiredQuantity"].(int); ok {
				n.RequiredQuantity = requiredQuantity
			}

			if reachedQuantity, ok := patch["reachedQuantity"].(int); ok {
				n.ReachedQuantity = reachedQuantity
			}

			withoutDueDate := patch["withoutDueDate"].(bool)
			if withoutDueDate {
				n.DueDate = nil
			}

			if dueDate, ok := patch["dueDate"].(time.Time); ok {
				if withoutDueDate {
					return nil, errors.New("parameters withoutDueDate and dueDate can't be used together")
				}
				n.DueDate = &dueDate
			}

			if categoryID, ok := patch["categoryId"].(int); ok {
				n.CategoryID = int64(categoryID)
			}

			if status, ok := patch["status"]; ok {
				switch status {
				case model.NeedStatusActive:
					n.Status = model.NeedStatusActive
				case model.NeedStatusInactive:
					n.Status = model.NeedStatusInactive
				default:
					return nil, errors.New("status should be active or inactive only")
				}
			}

			*n, err = update(*n)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{"need": n}, nil
		},
	}
}
