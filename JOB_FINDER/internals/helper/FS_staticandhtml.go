package helper

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func GetStatic(router *chi.Mux, zaplog *zap.SugaredLogger) {
	fs := http.FileServer(http.Dir("./web/static"))
	zaplog.Info("static")
	router.Handle(
		"/static/*",
		http.StripPrefix("/static/", fs),
	)
}

func GetJSScript(router *chi.Mux, zaplog *zap.SugaredLogger) {
	fs := http.FileServer(http.Dir("./web/scripts"))
	zaplog.Info("scripts")
	router.Handle(
		"/scripts/*",
		http.StripPrefix("/scripts/", fs),
	)
}
