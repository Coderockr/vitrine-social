package graphql

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

type (
	getCatFn              func(int64) (*model.Category, error)
	getCatByResolveParams func(graphql.ResolveParams) (*model.Category, error)
	getCatNeedsCountFn    func(*model.Category) (int64, error)
)

func newCategoryQuery(get getCatFn, count getCatNeedsCountFn) *graphql.Field {

	categoryType.AddFieldConfig("needsCount", &graphql.Field{
		Type: graphql.NewNonNull(graphql.Int),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if c, ok := p.Source.(*model.Category); ok && c != nil {
				return count(c)
			}

			if c, ok := p.Source.(model.Category); ok {
				return count(&c)
			}

			return 0, nil
		},
	})

	f := newCategoryField(func(p graphql.ResolveParams) (*model.Category, error) {
		if id, ok := p.Args["id"].(int); ok {
			return get(int64(id))
		}
		return nil, nil
	})

	f.Name = "CategoryQuery"
	f.Description = "Retrieves a Category by its ID"
	f.Args = graphql.FieldConfigArgument{
		"id": nonNullIntArgument,
	}
	return f
}

func newCategoryField(get getCatByResolveParams) *graphql.Field {
	return &graphql.Field{
		Type: categoryType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return get(p)
		},
	}
}
