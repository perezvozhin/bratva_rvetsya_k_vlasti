package api_

import (
	"fmt"
	"net/http"
)

func TestApi() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("TestApi")
	}
}
