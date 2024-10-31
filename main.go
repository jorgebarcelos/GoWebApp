package main

import ("github.com/jorgebarcelos/GoWebApp/endpoints"
		"github.com/jorgebarcelos/GoWebApp/server"
)

func main() {
	endpoints.Endpoints()
	server.ServerStart()
}