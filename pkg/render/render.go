package render

import (
	"bytes"
	"github.com/jorgebarcelos/GoWebApp/pkg/config"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var app *config.AppConfig

// sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var templateCache map[string]*template.Template
	
	if app.UseCache {
		// create a template cache
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	// get requested template from cache
	myTmpl, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	_ = myTmpl.Execute(buf, nil)

	//render template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all files named *-page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*-page.tmpl")
	if err != nil {
		return myCache, err
	}

	// loop trought all files ending with *-page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		layoutSet, err := filepath.Glob("./templates/*-layout.tmpl")

		if err != nil {
			return myCache, err
		}
		if len(layoutSet) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*-layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = templateSet
	}

	return myCache, nil
}
