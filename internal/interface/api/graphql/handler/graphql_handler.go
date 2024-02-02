package handler

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	gql "github.con/AndyGo/go-ddd/internal/interface/api/graphql/gql"
)

type GraphQLRequest struct {
	Query    string `json:"query"`
	Mutation string `json:"mutation"`
}

func GraphqlHandler(c echo.Context) error {
	var graphqlRequest GraphQLRequest
	if err := c.Bind(&graphqlRequest); err != nil {
		return c.JSON(http.StatusBadRequest, "Error al analizar la solicitud: "+err.Error())
	}

	var requestString string

	if graphqlRequest.Mutation == "" && graphqlRequest.Query == "" {
		return c.JSON(http.StatusBadRequest, "Se requiere un campo 'query' o 'mutation' en la solicitud.")
	}

	if graphqlRequest.Query != "" {
		requestString = graphqlRequest.Query
	}

	if graphqlRequest.Mutation != "" {
		requestString = graphqlRequest.Mutation
	}

	params := graphql.Params{
		Schema:        gql.Schema,
		RequestString: requestString,
	}

	result := graphql.Do(params)
	if len(result.Errors) > 0 {
		return c.JSON(http.StatusBadRequest, result.Errors)
	}

	return c.JSON(http.StatusOK, result)
}
