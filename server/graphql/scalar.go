package graphql

import (
	"log"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

const dateFormat = "2006-01-02"

func serializeDate(value interface{}) interface{} {
	switch value := value.(type) {
	case time.Time:
		return value.Format(dateFormat)
	case *time.Time:
		return serializeDate(*value)
	default:
		return nil
	}
}

func unserializeDate(value interface{}) interface{} {
	switch value := value.(type) {
	case string:
		log.Printf("string: %s", value)
		t, err := time.Parse(dateFormat, value)
		if err != nil {
			return nil
		}

		return t
	case []byte:
		return unserializeDate(string(value))
	case *string:
		return unserializeDate(string(*value))
	default:
		return nil
	}
}

var Date = graphql.NewScalar(graphql.ScalarConfig{
	Name: "Date",
	Description: "The `Date` scalar type represents a Date with the format \"yyyy-mm-dd\"." +
		" The Date is serialized as an string",
	Serialize:  serializeDate,
	ParseValue: unserializeDate,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return unserializeDate(valueAST.Value)
		}
		return nil
	},
})
