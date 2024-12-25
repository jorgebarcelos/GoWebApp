package handlers

import (
	"net/http"

	"github.com/jorgebarcelos/GoWebApp/pkg/config"
	"github.com/jorgebarcelos/GoWebApp/pkg/models"
	"github.com/jorgebarcelos/GoWebApp/pkg/render"
)

// is the repository type
type Repository struct {
	App *config.AppConfig
}

// repository used by the handlers
var Repo *Repository

//creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo = r
}

// home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home-page.tmpl", &models.TemplateData{})
}

// about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["Cú"] = "Hello Cú"

	render.RenderTemplate(w, "about-page.tmpl", &models.TemplateData{})
}