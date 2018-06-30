package graphql

import (
	"fmt"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

var (
	dateInput = &graphql.InputObjectFieldConfig{
		Type: Date,
	}

	intListArgument = &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.Int),
	}

	intArgument = &graphql.ArgumentConfig{
		Type: graphql.Int,
	}

	stringField = &graphql.Field{
		Type: graphql.String,
	}

	intField = &graphql.Field{
		Type: graphql.Int,
	}

	nonNullIntArgument = &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	}

	nonNullIntInput = &graphql.InputObjectFieldConfig{
		Type: graphql.NewNonNull(graphql.Int),
	}

	intInput = &graphql.InputObjectFieldConfig{
		Type: graphql.Int,
	}

	intListInput = &graphql.InputObjectFieldConfig{
		Type: graphql.NewList(graphql.NewNonNull(graphql.Int)),
	}

	nonNullIntField = &graphql.Field{
		Type: graphql.NewNonNull(graphql.Int),
	}

	stringArgument = &graphql.ArgumentConfig{
		Type: graphql.String,
	}

	nonNullStringInput = &graphql.InputObjectFieldConfig{
		Type: graphql.NewNonNull(graphql.String),
	}

	stringInput = &graphql.InputObjectFieldConfig{
		Type: graphql.String,
	}

	nonNullStringArgument = &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	}

	nonNullStringField = &graphql.Field{
		Type: graphql.NewNonNull(graphql.String),
	}

	organizationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Organization",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if o, ok := p.Source.(*model.Organization); ok && o != nil {
						return o.ID, nil
					}
					return nil, nil
				},
			},
			"images": &graphql.Field{
				Type: graphql.NewList(organizationImageType),
			},
			"name": nonNullStringField,
			"logo": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if o, ok := p.Source.(*model.Organization); ok && o != nil && o.Logo != nil {
						return o.Logo.URL, nil
					}
					return nil, nil
				},
			},
			"slug":  nonNullStringField,
			"phone": nonNullStringField,
			"about": nonNullStringField,
			"video": nonNullStringField,
			"email": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if o, ok := p.Source.(*model.Organization); ok && o != nil {
						return o.Email, nil
					}
					return nil, nil
				},
			},
			"address": &graphql.Field{
				Type: graphql.NewNonNull(addressType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if o, ok := p.Source.(*model.Organization); ok && o != nil {
						return o.Address, nil
					}
					return nil, nil
				},
			},
		},
	})

	addressType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Address",
		Fields: graphql.Fields{
			"street":        stringField,
			"number":        stringField,
			"complement":    stringField,
			"neighbordhood": stringField,
			"city":          stringField,
			"state":         stringField,
			"zipcode":       stringField,
		},
	})

	organizationImageType = graphql.NewObject(graphql.ObjectConfig{
		Name: "OrganizationImage",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if i, ok := p.Source.(model.OrganizationImage); ok {
						return i.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if i, ok := p.Source.(model.OrganizationImage); ok {
						return i.Name, nil
					}
					return nil, nil
				},
			},
			"url": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if i, ok := p.Source.(model.OrganizationImage); ok {
						return i.URL, nil
					}
					return nil, nil
				},
			},
		},
	})

	needStatusEnum = graphql.NewEnum(graphql.EnumConfig{
		Name:        "NeedStatus",
		Description: "Status of a Need",
		Values: graphql.EnumValueConfigMap{
			"ACTIVE": &graphql.EnumValueConfig{
				Value:       model.NeedStatusActive,
				Description: "A active Need",
			},
			"INACTIVE": &graphql.EnumValueConfig{
				Value:       model.NeedStatusInactive,
				Description: "A inactive Need",
			},
		},
	})

	needStatusInput = &graphql.InputObjectFieldConfig{
		Type: needStatusEnum,
	}

	needImageType = graphql.NewObject(graphql.ObjectConfig{
		Name: "NeedImage",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if i, ok := p.Source.(model.NeedImage); ok {
						return i.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if i, ok := p.Source.(model.NeedImage); ok {
						return i.Name, nil
					}
					return nil, nil
				},
			},
			"url": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if i, ok := p.Source.(model.NeedImage); ok {
						return i.URL, nil
					}
					return nil, nil
				},
			},
		},
	})

	needType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Need",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if n, ok := p.Source.(*model.Need); ok {
						return n.ID, nil
					}

					if s, ok := p.Source.(model.SearchNeed); ok {
						return s.ID, nil
					}

					return nil, fmt.Errorf("invalid source")
				},
			},
			"title": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if n, ok := p.Source.(*model.Need); ok {
						return n.Title, nil
					}

					if s, ok := p.Source.(model.SearchNeed); ok {
						return s.Title, nil
					}

					return nil, fmt.Errorf("invalid source")
				},
			},
			"description": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if n, ok := p.Source.(*model.Need); ok {
						return n.Description, nil
					}

					if s, ok := p.Source.(model.SearchNeed); ok {
						return s.Description, nil
					}

					return nil, fmt.Errorf("invalid source")
				},
			},
			"requiredQuantity": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if n, ok := p.Source.(*model.Need); ok {
						return n.RequiredQuantity, nil
					}

					if s, ok := p.Source.(model.SearchNeed); ok {
						return s.RequiredQuantity, nil
					}

					return nil, fmt.Errorf("invalid source")
				},
			},

			"reachedQuantity": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if n, ok := p.Source.(*model.Need); ok {
						return n.ReachedQuantity, nil
					}

					if s, ok := p.Source.(model.SearchNeed); ok {
						return s.ReachedQuantity, nil
					}

					return nil, fmt.Errorf("invalid source")
				},
			},
			"unit": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if n, ok := p.Source.(*model.Need); ok {
						return n.Unit, nil
					}

					if s, ok := p.Source.(model.SearchNeed); ok {
						return s.Unit, nil
					}

					return nil, fmt.Errorf("invalid source")
				},
			},
			"dueDate": &graphql.Field{
				Type: Date,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if n, ok := p.Source.(*model.Need); ok {
						return n.DueDate, nil
					}

					if s, ok := p.Source.(model.SearchNeed); ok {
						return s.DueDate, nil
					}

					return nil, fmt.Errorf("invalid source")
				},
			},
			"status": &graphql.Field{
				Type: needStatusEnum,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if n, ok := p.Source.(*model.Need); ok {
						return n.Status, nil
					}

					if s, ok := p.Source.(model.SearchNeed); ok {
						return s.Status, nil
					}

					return nil, fmt.Errorf("invalid source")
				},
			},
			"createdAt": &graphql.Field{
				Type: graphql.NewNonNull(DateTime),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if n, ok := p.Source.(*model.Need); ok {
						return n.CreatedAt, nil
					}

					if s, ok := p.Source.(model.SearchNeed); ok {
						return s.CreatedAt, nil
					}

					return nil, fmt.Errorf("invalid source")
				},
			},
			"updatedAt": &graphql.Field{
				Type: DateTime,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if n, ok := p.Source.(*model.Need); ok {
						return n.UpdatedAt, nil
					}

					if s, ok := p.Source.(model.SearchNeed); ok {
						return s.UpdatedAt, nil
					}

					return nil, fmt.Errorf("invalid source")
				},
			},
			"images": &graphql.Field{
				Type: graphql.NewList(needImageType),
			},
		},
	})

	categoryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			"id":   nonNullIntField,
			"name": nonNullStringField,
			"slug": nonNullStringField,
		},
	})
)
