package main

import (
	"log"

	"github.com/jorgebarcelos/GoWebApp/pkg/config"
	"github.com/jorgebarcelos/GoWebApp/pkg/endpoints"
	"github.com/jorgebarcelos/GoWebApp/pkg/handlers"
	"github.com/jorgebarcelos/GoWebApp/pkg/render"
	"github.com/jorgebarcelos/GoWebApp/pkg/server"
)

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	endpoints.Endpoints()
	server.ServerStart()
}