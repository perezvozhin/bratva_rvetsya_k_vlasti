package main

import (
	loggersystem "JOB_FINDER/internals/logger"
	"fmt"
	"net/http"
)

func main() {

	//инит вспом штук
	logger := loggersystem.Init()

	//инит сервера
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "yo world")
	})
	//просто текст для бренча
	fmt.Println("starting server :8080")

	logger.Info("started server")
	http.ListenAndServe(":8080", nil)
}
