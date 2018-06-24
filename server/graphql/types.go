package graphql

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gobuffalo/pop/nulls"
	"github.com/graphql-go/graphql"
)

var stringArgument = &graphql.ArgumentConfig{
	Type: graphql.String,
}

var intListArgument = &graphql.ArgumentConfig{
	Type: graphql.NewList(graphql.Int),
}

var intArgument = &graphql.ArgumentConfig{
	Type: graphql.Int,
}

var stringField = &graphql.Field{
	Type: graphql.String,
}
var intField = &graphql.Field{
	Type: graphql.Int,
}
var nonNullIntArgument = &graphql.ArgumentConfig{
	Type: graphql.NewNonNull(graphql.Int),
}

var nonNullStringArgument = &graphql.ArgumentConfig{
	Type: graphql.NewNonNull(graphql.String),
}

var nonNullStringField = &graphql.Field{
	Type: graphql.NewNonNull(graphql.String),
}
var nonNullIntField = &graphql.Field{
	Type: graphql.NewNonNull(graphql.Int),
}

var organizationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Organization",
	Fields: graphql.Fields{
		"id":    nonNullIntField,
		"name":  nonNullStringField,
		"logo":  nonNullStringField,
		"slug":  nonNullStringField,
		"phone": nonNullStringField,
		"about": nonNullStringField,
		"video": nonNullStringField,
		"email": nonNullStringField,
		"address": &graphql.Field{
			Type: addressType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if o, ok := p.Source.(organizationJSON); ok {
					return o.Address, nil
				}
				return nil, nil
			},
		},
	},
})

var addressType = graphql.NewObject(graphql.ObjectConfig{
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

type addressJSON struct {
	Street       string       `json:"street"`
	Number       string       `json:"number"`
	Complement   nulls.String `json:"complement"`
	Neighborhood string       `json:"neighborhood"`
	City         string       `json:"city"`
	State        string       `json:"state"`
	Zipcode      string       `json:"zipcode"`
}

type organizationJSON struct {
	ID      int64       `json:"id"`
	Name    string      `json:"name"`
	Logo    string      `json:"logo"`
	Slug    string      `json:"slug"`
	Phone   string      `json:"phone"`
	Video   string      `json:"video"`
	About   string      `json:"about"`
	Email   string      `json:"email"`
	Address addressJSON `json:"address"`
}

func orgToJSON(o *model.Organization) *organizationJSON {
	return &organizationJSON{
		ID:    o.ID,
		Name:  o.Name,
		Logo:  o.Logo.URL,
		Slug:  o.Slug,
		Phone: o.Phone,
		Video: o.Video,
		About: o.About,
		Email: o.Email,
		Address: addressJSON{
			Street:       o.Address.Street,
			Number:       o.Address.Number,
			Complement:   o.Address.Complement,
			Neighborhood: o.Address.Neighborhood,
			City:         o.Address.City,
			State:        o.Address.State,
			Zipcode:      o.Address.Zipcode,
		},
	}
}

var needStatusEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "NeedStatus",
	Description: "Status of a Need",
	Values: graphql.EnumValueConfigMap{
		"ACTIVE": &graphql.EnumValueConfig{
			Value:       "ACTIVE",
			Description: "A active Need",
		},
		"INACTIVE": &graphql.EnumValueConfig{
			Value:       "INACTIVE",
			Description: "A inactive Need",
		},
	},
})

var needType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Need",
	Fields: graphql.Fields{
		"id":               nonNullIntField,
		"title":            nonNullStringField,
		"description":      nonNullStringField,
		"requiredQuantity": intField,
		"reachedQuantity":  intField,
		"unity":            stringField,
		"status": &graphql.Field{
			Type: needStatusEnum,
		},
	},
})

var categoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Category",
	Fields: graphql.Fields{
		"id":   nonNullIntField,
		"name": nonNullStringField,
		"slug": nonNullStringField,
	},
})
