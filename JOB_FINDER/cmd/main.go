package main

import (
	api_ "JOB_FINDER/api/rest"
	"JOB_FINDER/api/view"
	"JOB_FINDER/httpmw"
	"JOB_FINDER/internals/FS_config"
	"JOB_FINDER/internals/helper"
	loggersystem "JOB_FINDER/internals/logger"
	"JOB_FINDER/usr_service"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// писька тряс
func main() {
	//инициализация чи
	router := chi.NewMux()

	//инит вспом штук
	logger := loggersystem.Init()
	cfg := FS_config.Init(logger)
	CV := helper.CheckDirectoryForCV(cfg.PathFilesystem, logger)
	//init services

	CS_userService := usr_service.Init(logger, CV)

	//временно
	fmt.Println(CS_userService)
	//статика (html + css)
	helper.GetStatic(router, logger)
	//статика (js)
	helper.GetJSScript(router, logger)
	//парсер html
	tmplparser := template.Must(template.ParseGlob("web/*.html"))
	//инит сервисов
	// - userService := usr_service.Init(nil)

	router.Get("/insert", api_.InsertApi(tmplparser))

	//коллекция api
	router.Group(func(r chi.Router) {
		//апи для ввода ключа к гемини (триггерится в случае, если строка конфига пуста)
		//смотри логику в httpmw.ApiCheckMiddleWare(cfg.ApiKey)
		r.Use(httpmw.ApiCheckMiddleWare(cfg.ApiKey))
		r.Handle("/", view.MainPageDrawer(tmplparser))

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

// вова
