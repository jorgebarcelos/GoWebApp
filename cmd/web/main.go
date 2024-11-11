package main

import (
	"github.com/jorgebarcelos/GoWebApp/pkg/cache"
	"github.com/jorgebarcelos/GoWebApp/pkg/endpoints"
	"github.com/jorgebarcelos/GoWebApp/pkg/server"
)

func main() {

	cache.UseCache()
	endpoints.Endpoints()
	server.ServerStart()
}