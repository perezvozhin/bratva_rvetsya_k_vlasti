package main

import (
	api_ "JOB_FINDER/api"
	"JOB_FINDER/httpmw"
	"JOB_FINDER/internals/FS_config"
	loggersystem "JOB_FINDER/internals/logger"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// ad ad a da a d
func main() {

	//инит вспом штук
	logger := loggersystem.Init()
	cfg := FS_config.Init(logger)

	//инициализация чи
	router := chi.NewMux()

	//глобальный миддлвар на проверку api gemini
	router.Use(httpmw.ApiCheckMiddleWare(cfg.ApiKey))

	//коллекция api
	router.Route("/api", func(r chi.Router) {
		//апи для ввода ключа к гемини (триггерится в случае, если строка конфига пуста)
		//смотри логику в httpmw.ApiCheckMiddleWare(cfg.ApiKey)

		router.Get("/insert", api_.InsertApi(cfg.PathStatic))
	})

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
