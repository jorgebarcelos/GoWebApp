package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)


func RenderTemplate(w http.ResponseWriter, tmpl string){
	// create a template cache
	templateCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get requested template from cache
	myTmpl, ok := templateCache[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = myTmpl.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	//render template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string] *template.Template, error) {
	myCache := map[string] *template.Template{}

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