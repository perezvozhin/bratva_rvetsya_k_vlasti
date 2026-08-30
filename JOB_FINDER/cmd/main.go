package main

import (
	"JOB_FINDER/api"
	"JOB_FINDER/caller"
	"JOB_FINDER/gem_service"
	"JOB_FINDER/httpmw"
	config "JOB_FINDER/internals/FS_config"
	loggersystem "JOB_FINDER/internals/logger"
	"JOB_FINDER/storage"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewMux()

	logger := loggersystem.Init()

	cfg := config.NewConfigMust()

	mcp := caller.NewCaller(logger)
	gemService := gem_service.NewGeminiService(context.Background(), mcp, cfg.APIKey, cfg.PathToChats, cfg.DefaultModel)

	store, err := storage.NewStore(cfg.PathToInterview, logger)
	if err != nil {
		logger.Fatal("failed to init storage", "err", err)
		panic(err)
	}

	chatHandler := api.NewChatHandler(gemService, store, logger)

	// repo, err := sqlite.NewVacancyRepo(cfg.PathToDB)
	// if err != nil {
	// 	logger.Fatal("failed to init vacancy repo", "err", err)
	// 	panic(err)
	// }
	// defer repo.Close()
	//
	// parserService := parser.NewParserService(repo, logger,
	// 	providers.NewHabr(mcp),
	// 	providers.NewHH(mcp, cfg.HHUserAgent),
	// )
	// vacancyHandler := api.NewVacancyHandler(parserService, logger)

	router.Group(func(r chi.Router) {
		// апи для ввода ключа к гемини
		r.Use(httpmw.ApiCheckMiddleWare(cfg.APIKey))
		chatHandler.RegisterRoutes(r)
		// vacancyHandler.RegisterRoutes(r)
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
