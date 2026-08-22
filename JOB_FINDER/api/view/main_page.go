package view

import (
	"html/template"
	"net/http"
)

func MainPageDrawer(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tmpl.ExecuteTemplate(w, "index.html", nil)
	}
}
