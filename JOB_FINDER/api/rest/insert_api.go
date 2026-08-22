package rest

import (
	"html/template"
	"net/http"
)

func InsertApi(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tmpl.ExecuteTemplate(w, "api_insert_page.html", nil)
	}
}
