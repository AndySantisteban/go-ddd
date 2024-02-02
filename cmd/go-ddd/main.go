package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.con/AndyGo/go-ddd/internal/application/services"
	mssql "github.con/AndyGo/go-ddd/internal/infrastructure/db/mssql/respoitories"
	gqlHandler "github.con/AndyGo/go-ddd/internal/interface/api/graphql/handler"
	"github.con/AndyGo/go-ddd/internal/interface/api/rest"
	web "github.con/AndyGo/go-ddd/internal/interface/web"
	"github.con/AndyGo/go-ddd/internal/interface/web/react"
)

func main() {
	port := ":8080"
	if err := web.Execute("cmd/go-app"); err != nil {
		fmt.Println("Error:", err)
	}

	// Instance Services
	partnerRepo := mssql.NewPLSPartnerRepository()
	partnerService := services.NewPartnerService(partnerRepo)
	allnoteRepo := mssql.NewALLNoteRepository()
	allnoteService := services.NewALLNoteService(allnoteRepo)
	e := echo.New()

	react.RegisterHandlersWebApp(e)
	e.Use(middleware.Logger(), middleware.Recover())

	rest.NewPartnerController(e, partnerService)
	rest.NewALLNoteController(e, allnoteService)
	e.Use(echoprometheus.NewMiddleware("myapp"))
	e.GET("/metrics", echoprometheus.NewHandler())
	e.POST("/graphql", gqlHandler.GraphqlHandler)
	e.File("/graphiql", "internal/interface/api/graphql/public/playground.html")

	if err := e.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
