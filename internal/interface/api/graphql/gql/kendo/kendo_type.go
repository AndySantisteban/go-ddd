package kendo

import "github.com/graphql-go/graphql"

func NewResponseKendoMvcType(dataType graphql.Output) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "DataSourceRequestResponse",
			Fields: graphql.Fields{
				"Data": &graphql.Field{
					Type: graphql.NewList(dataType),
				},
				"Total": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
}
