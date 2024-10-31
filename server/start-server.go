package server

import (
	"fmt"
	"net/http"
)

const portNumber = "localhost:8000"


func ServerStart(){
	fmt.Printf("Start at port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
