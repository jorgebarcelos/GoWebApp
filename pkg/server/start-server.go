package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jorgebarcelos/GoWebApp/pkg/config"
	"github.com/jorgebarcelos/GoWebApp/pkg/handlers"
	"github.com/jorgebarcelos/GoWebApp/pkg/render"
	"github.com/jorgebarcelos/GoWebApp/pkg/routes"
)

const portNumber = "localhost:8000"
var app config.AppConfig

func ServerStart(){
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	fmt.Printf("Start at port %s", portNumber)
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes.Routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
