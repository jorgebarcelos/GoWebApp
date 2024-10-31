package endpoints

import ("net/http"
		"github.com/jorgebarcelos/GoWebApp/pkg/handlers"
)


func Endpoints() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
}