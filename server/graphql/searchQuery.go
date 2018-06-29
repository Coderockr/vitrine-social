package graphql

import (
	"math"

	"github.com/Coderockr/vitrine-social/server/db/repo"
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

var (
	orderEnum = graphql.NewEnum(graphql.EnumConfig{
		Name:        "SearchOrder",
		Description: "Order will be accending or descending",
		Values: graphql.EnumValueConfigMap{
			"ASC": &graphql.EnumValueConfig{
				Value: "asc",
			},
			"DESC": &graphql.EnumValueConfig{
				Value: "desc",
			},
		},
	})

	orderByEnum = graphql.NewEnum(graphql.EnumConfig{
		Name:        "SearchOrderBy",
		Description: "Which kind of ordering will be used",
		Values: graphql.EnumValueConfigMap{
			"ID": &graphql.EnumValueConfig{
				Value:       "id",
				Description: "Order by Need's ID",
			},
			"UPDATED_AT": &graphql.EnumValueConfig{
				Value:       "updated_at",
				Description: "Order by Last Need's Update date",
			},
			"CREATED_AT": &graphql.EnumValueConfig{
				Value:       "created_at",
				Description: "Order by Need's Creation date",
			},
		},
	})

	searchInput = graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "SearchNeedsInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"text":           stringInput,
			"categories":     intListInput,
			"organizationId": intInput,
			"orderBy": &graphql.InputObjectFieldConfig{
				Type:         orderByEnum,
				DefaultValue: "created_at",
			},
			"order": &graphql.InputObjectFieldConfig{
				Type:         orderEnum,
				DefaultValue: "desc",
			},
			"page": &graphql.InputObjectFieldConfig{
				Type:         graphql.Int,
				DefaultValue: 1,
			},
		},
	})
)

func newSearchQuery(search searchNeedsFn) *graphql.Field {
	f := newSearchNeedField(
		search,
		func(p graphql.ResolveParams) (searchParams, error) {
			sp := searchParams{}

			input, ok := p.Args["input"].(map[string]interface{})
			if !ok {
				return sp, nil
			}

			sp.Text, _ = input["text"].(string)
			sp.Categories, _ = input["categories"].([]int)
			id, _ := input["organizationId"].(int)
			sp.OrganizationID = int64(id)
			sp.OrderBy, _ = input["orderBy"].(string)
			sp.Order, _ = input["order"].(string)
			sp.Page, _ = input["page"].(int)

			return sp, nil
		},
		graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: searchInput,
			},
		},
	)
	f.Name = "SearchQuery"
	f.Description = "Search active needs on the database"

	return f
}

type (
	searchNeedsFn      func(text string, categoriesID []int, organizationsID int64, status string, orderBy string, order string, page int) (results []model.SearchNeed, count int, err error)
	parseSearchInputFn func(graphql.ResolveParams) (searchParams, error)

	searchParams struct {
		Text           string
		Categories     []int
		OrganizationID int64
		Status         string
		OrderBy        string
		Order          string
		Page           int
	}
)

var (
	pageInfoType = graphql.NewObject(graphql.ObjectConfig{
		Name: "PageInfo",
		Fields: graphql.Fields{
			"totalResults": nonNullIntField,
			"totalPages":   nonNullIntField,
			"currentPage":  nonNullIntField,
		},
	})

	paginatedNeedsPayload = graphql.NewObject(graphql.ObjectConfig{
		Name: "PaginatedNeedsPayload",
		Fields: graphql.Fields{
			"results": &graphql.Field{
				Type: graphql.NewList(graphql.NewNonNull(needType)),
			},
			"pageInfo": &graphql.Field{
				Type: pageInfoType,
			},
		},
	})
)

func newSearchNeedField(search searchNeedsFn, parseInput parseSearchInputFn, args graphql.FieldConfigArgument) *graphql.Field {
	return &graphql.Field{
		Args: args,
		Type: paginatedNeedsPayload,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			sp, err := parseInput(p)
			if err != nil {
				return nil, err
			}
			if sp.Page == 0 {
				sp.Page = 1
			}

			rs, c, err := search(
				sp.Text,
				sp.Categories,
				sp.OrganizationID,
				sp.Status,
				sp.OrderBy,
				sp.Order,
				sp.Page,
			)
			if err != nil {
				return nil, err
			}

			payload := map[string]interface{}{
				"results": rs,
				"pageInfo": map[string]interface{}{
					"totalResults": c,
					"totalPages":   int(math.Ceil(float64(c) / repo.ResultsPerPage)),
					"currentPage":  sp.Page,
				},
			}

			return payload, nil
		},
	}
}
