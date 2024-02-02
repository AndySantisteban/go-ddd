package gql

import (
	"github.com/graphql-go/graphql"
	AllNoteGQL "github.con/AndyGo/go-ddd/internal/interface/api/graphql/gql/allnote"
)

var RootQueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"ListALLNote": AllNoteGQL.ALLNoteRequest,
		},
	},
)

var RootMutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"AddNote": AllNoteGQL.CreateALLNoteRequest,
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    RootQueryType,
		Mutation: RootMutationType,
	},
)
