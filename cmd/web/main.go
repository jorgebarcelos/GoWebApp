package main

import ("github.com/jorgebarcelos/GoWebApp/pkg/endpoints"
		"github.com/jorgebarcelos/GoWebApp/pkg/server"
)

func main() {
	endpoints.Endpoints()
	server.ServerStart()
}