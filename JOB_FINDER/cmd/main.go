package main

import (
	api_ "JOB_FINDER/api"
	"JOB_FINDER/internals/FS_config"
	loggersystem "JOB_FINDER/internals/logger"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	//инит вспом штук
	logger := loggersystem.Init()
	cfg := FS_config.Init(logger)

	//инициализация чи
	router := chi.NewMux()

	//гет запрос с роутом
	router.Get("/test", api_.TestApi())

	//добавим чи, как мультиплексер
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}
	logger.Info("started server")
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatal("error starting server", "error", err)
		panic(err)
	}
}
