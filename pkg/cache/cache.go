package cache

import (
	"log"

	"github.com/jorgebarcelos/GoWebApp/pkg/config"
	"github.com/jorgebarcelos/GoWebApp/pkg/handlers"
	"github.com/jorgebarcelos/GoWebApp/pkg/render"
)

var app config.AppConfig

func UseCache(){
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
}