package allnote

import (
	"github.com/graphql-go/graphql"
	"github.con/AndyGo/go-ddd/internal/interface/api/graphql/gql/kendo"
	resolver "github.con/AndyGo/go-ddd/internal/interface/api/graphql/resolvers"
)

// RESPONSE TYPES
var ALLNoteType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AllNote",
		Fields: graphql.Fields{
			"Uid":                 &graphql.Field{Type: graphql.Int},
			"ParentUid":           &graphql.Field{Type: graphql.String},
			"UserUid":             &graphql.Field{Type: graphql.String},
			"SubjectOld":          &graphql.Field{Type: graphql.String},
			"CategoryUid":         &graphql.Field{Type: graphql.Int},
			"SubCategory":         &graphql.Field{Type: graphql.Int},
			"Note":                &graphql.Field{Type: graphql.String},
			"Note_Plain":          &graphql.Field{Type: graphql.String},
			"Priority":            &graphql.Field{Type: graphql.Int},
			"CommunicationCode":   &graphql.Field{Type: graphql.Int},
			"IsReminder":          &graphql.Field{Type: graphql.Boolean},
			"CompletedBy":         &graphql.Field{Type: graphql.String},
			"CompletedDate":       &graphql.Field{Type: graphql.DateTime},
			"ScheduledStart":      &graphql.Field{Type: graphql.DateTime},
			"ScheduledEnd":        &graphql.Field{Type: graphql.DateTime},
			"OldRecId":            &graphql.Field{Type: graphql.String},
			"AppTimeStamp":        &graphql.Field{Type: graphql.DateTime},
			"AppLastUpdatedBy":    &graphql.Field{Type: graphql.String},
			"AppCreatedBy":        &graphql.Field{Type: graphql.Int},
			"ActivityUid":         &graphql.Field{Type: graphql.String},
			"Subject":             &graphql.Field{Type: graphql.String},
			"ContactName":         &graphql.Field{Type: graphql.String},
			"AppTimeStampYear":    &graphql.Field{Type: graphql.Int},
			"SysTimeStamp":        &graphql.Field{Type: graphql.String},
			"Note_PlainResume":    &graphql.Field{Type: graphql.String},
			"ParentDepartmentUid": &graphql.Field{Type: graphql.Int},
			"ParentAccount":       &graphql.Field{Type: graphql.String},
			"ParentType":          &graphql.Field{Type: graphql.Int},
			"AppLastUpdated":      &graphql.Field{Type: graphql.DateTime},
		},
	},
)

var ALLNoteCreateType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CreateALLNote",
		Fields: graphql.Fields{
			"Uid": &graphql.Field{Type: graphql.Int},
		},
	},
)

// ARGS VALUES

var ALLNoteResposeArgs = graphql.FieldConfigArgument{
	"dataSourceRequest": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var CreateALLNoteArgs = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "AllNoteInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"ParentUid": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"UserUid": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"SubjectOld": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"CategoryUid": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"SubCategory": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"Note": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"Note_Plain": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"Priority": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"CommunicationCode": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"IsReminder": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
			"CompletedBy": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"CompletedDate": &graphql.InputObjectFieldConfig{
				Type: graphql.DateTime,
			},
			"ScheduledStart": &graphql.InputObjectFieldConfig{
				Type: graphql.DateTime,
			},
			"ScheduledEnd": &graphql.InputObjectFieldConfig{
				Type: graphql.DateTime,
			},
			"OldRecId": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"AppLastUpdatedBy": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"AppCreatedBy": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"ActivityUid": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Subject": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"ContactName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"AppTimeStampYear": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"Note_PlainResume": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"SysTimeStamp": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"ParentDepartmentUid": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"ParentAccount": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"ParentType": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"AppLastUpdated": &graphql.InputObjectFieldConfig{
				Type: graphql.DateTime,
			},
		},
	},
)
var CreateALLNoteResposeArgs = graphql.FieldConfigArgument{
	"note": &graphql.ArgumentConfig{
		Type: CreateALLNoteArgs,
	},
}

// QUERY OR MUTATIONS CONFIG

var ALLNoteRequest = &graphql.Field{
	Type:    kendo.NewResponseKendoMvcType(ALLNoteType),
	Args:    ALLNoteResposeArgs,
	Resolve: resolver.GetAllNotesGQL,
}

var CreateALLNoteRequest = &graphql.Field{
	Type:    ALLNoteCreateType,
	Args:    CreateALLNoteResposeArgs,
	Resolve: resolver.CreateAllNotesGQL,
}
