package view

import (
	"JOB_FINDER/gem_service"
	"html/template"
	"net/http"
)

type VI_TextResponse struct {
	text string
}

func MainPageDrawer(gemservice *gem_service.GeminiService, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vistrct VI_TextResponse
		ch := gemservice.AskStream(r.Context(), chat, text)

		go func() {
			select {
			case <-ch:
				text <- ch
			case <-r.Context().Done():
				chunkerr
			}
		}()

		tmpl.ExecuteTemplate(w, "index.html", vistrct)
	}
}
