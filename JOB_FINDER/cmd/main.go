package main

import (
	"JOB_FINDER/internals/FS_config"
	loggersystem "JOB_FINDER/internals/logger"
	"fmt"
	"net/http"
)

func main() {

	//инит вспом штук
	logger := loggersystem.Init()
	cfg := FS_config.Init(logger)

	//инит сервера
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "yo world")
	})
	//просто текст для бренча
	fmt.Println("starting server :8080")

	logger.Info("started server")
	http.ListenAndServe(cfg.Port, nil)
}
