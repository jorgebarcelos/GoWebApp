package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello Motherfucker")
		if err != nil{
			fmt.Println(err)
		}
		fmt.Printf("number of bytes: %d\n", n)
	})

	http.ListenAndServe(":8000", nil)
}