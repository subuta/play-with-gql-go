package main

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/graphql-go/graphql"
	"github.com/k0kubun/pp"
	"github.com/subuta/play-with-gql-go/internal/db"
	"log"
)

func main () {
	// ----- Try [graphql-go/graphql: An implementation of GraphQL for Go / Golang](https://github.com/graphql-go/graphql)
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			hello
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	//rJSON, _ := json.Marshal(r)

	pp.Println(r.Data)

	// ----- Try [Masterminds/squirrel: Fluent SQL generation for golang](https://github.com/Masterminds/squirrel)

	users := sq.Select("*").From("users").Join("emails USING (email_id)")

	active := users.Where(sq.Eq{"deleted_at": nil})

	sql_str, args, _ := active.ToSql()

	pp.Println(sql_str, args)

	// ----- Try DB operation with WASM fallback.
	db.Run()
}
