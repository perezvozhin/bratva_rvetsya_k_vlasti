package main

import (
	"JOB_FINDER/caller"
	"JOB_FINDER/gem_service"
	"JOB_FINDER/httpmw"
	config "JOB_FINDER/internals/FS_config"
	loggersystem "JOB_FINDER/internals/logger"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewMux()

	logger := loggersystem.Init()

	cfg := config.NewConfigMust()

	mcp := caller.Init(logger, "path")
	_ = gem_service.Init(context.Background(), mcp, cfg.APIKey)

	router.Group(func(r chi.Router) {
		// апи для ввода ключа к гемини
		r.Use(httpmw.ApiCheckMiddleWare(cfg.APIKey))
	})

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	logger.Info("started server at :", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal("error starting server", "error", err)
		panic(err)
	}
}
