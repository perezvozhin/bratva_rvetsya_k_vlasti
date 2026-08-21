package httpmw

import (
	"net/http"
)

func ApiCheckMiddleWare(api string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if api == "" {
				//роутим на страницу ввода api
				http.Redirect(w, r, "/api/insert", http.StatusFound)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
