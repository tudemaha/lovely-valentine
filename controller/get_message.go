package controller

import (
	"fmt"
	"net/http"
)

func GetIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		http.Redirect(w, r, "/create", http.StatusSeeOther)
	}
}
