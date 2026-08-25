package main

import (
	api_ "JOB_FINDER/api/rest"
	"JOB_FINDER/api/view"
	"JOB_FINDER/caller"
	"JOB_FINDER/gem_service"
	"JOB_FINDER/httpmw"
	"JOB_FINDER/internals/FS_config"
	"JOB_FINDER/internals/helper"
	loggersystem "JOB_FINDER/internals/logger"
	"JOB_FINDER/internals/repository/sqlite"
	"JOB_FINDER/usr_service"
	"JOB_FINDER/usr_service/parser"
	"context"
	"html/template"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// писька тряс
func main() {
	//инициализация чи
	router := chi.NewMux()

	//инит вспом штук
	logger := loggersystem.Init()
	mcpCaller := caller.Init(logger, "path")
	cfg := FS_config.Init()
	CS_database := sqlite.NewVacancyRepo("database FIX")
	CV := helper.CheckDirectoryForCV(cfg.PathFilesystem, logger)
	//init services
	CS_parse := parser.NewParserService(mcpCaller, CS_database)
	//зафикс + зафикс
	CS_userService := usr_service.Init(logger, CV, CS_parse, mcpCaller)

	MCP := caller.Init(logger, "path")

	CS_gemService := gem_service.Init(context.Background(), MCP, cfg.ApiKey)
	//временно

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
		r.Handle("/", view.MainPageDrawer(CS_gemService, tmplparser))
		//инициализируем апи для парсинга
		r.Handle()
		r.Handle("/parse", api_.ParseApi(CS_userService))
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

	go func() {
		for {
			cfg := FS_config.Init()
			cfgTime, _ := time.Parse(time.RFC3339, cfg.LastUpdate)
			now := time.Now()

			if now.Before(cfgTime.AddDate(0, 0, 7)) {
				time.Sleep(1 * time.Hour)
				continue
			}
			logger.Info("Прошло 7 дней")
			usr_service.Synctime(logger)

			time.Sleep(1 * time.Hour)
		}
	}()
}

// вова
