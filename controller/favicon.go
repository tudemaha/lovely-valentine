package controller

import "net/http"

func FaviconHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/public/favicon.ico")
	}
}
