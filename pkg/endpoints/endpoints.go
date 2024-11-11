package endpoints

import ("net/http"
		"github.com/jorgebarcelos/GoWebApp/pkg/handlers"
)


func Endpoints() {
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
}