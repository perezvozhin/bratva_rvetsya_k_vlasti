package main

import (
	"fmt"
	"net/http"
)

func main() {

	//инит сервера
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "yo world")
	})

	fmt.Println("starting server :8080")
	http.ListenAndServe(":8080", nil)
}
