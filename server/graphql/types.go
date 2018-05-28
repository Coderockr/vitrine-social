package graphql

import (
	"github.com/Coderockr/vitrine-social/server/model"
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

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"username": nonNullStringField,
		"password": nonNullStringField,
	},
})

var organizationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Organization",
	Fields: graphql.Fields{
		"id":      nonNullIntField,
		"name":    nonNullStringField,
		"logo":    nonNullStringField,
		"slug":    nonNullStringField,
		"address": nonNullStringField,
		"phone":   nonNullStringField,
		"resume":  nonNullStringField,
		"video":   nonNullStringField,
		"email":   nonNullStringField,
	},
})

type organizationJSON struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Logo    string `json:"logo"`
	Slug    string `json:"slug"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Video   string `json:"video"`
	Resume  string `json:"resume"`
	Email   string `json:"email"`
}

func orgToJSON(o *model.Organization) *organizationJSON {
	return &organizationJSON{
		ID:      o.ID,
		Name:    o.Name,
		Logo:    o.Logo,
		Slug:    o.Slug,
		Address: o.Address,
		Phone:   o.Phone,
		Video:   o.Video,
		Resume:  o.Resume,
		Email:   o.Email,
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
		"icon": nonNullStringField,
	},
})
