package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string){
	ParsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := ParsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home-page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about-page.tmpl")
}