package controller

import (
	"log"
	"net/http"
	"strings"
	"text/template"
	"valentine-quote/model"
)

func CheckPasswordHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paths := strings.Split(r.URL.Path, "/")
		var message model.Message

		if len(paths) < 3 {
			http.Redirect(w, r, "/create", http.StatusSeeOther)
			return
		}
		if len(paths) > 3 {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		id := paths[2]
		if id == "" {
			http.Redirect(w, r, "/create", http.StatusSeeOther)
			return
		}

		message, available := model.CheckAvailableID(id)
		if !available {
			http.Redirect(w, r, "/create", http.StatusSeeOther)
			return
		}

		data := struct {
			ID    string
			Magic string
		}{
			ID:    message.ID.Hex(),
			Magic: message.Magic,
		}

		t, err := template.ParseFiles("./view/password.html")
		if err != nil {
			log.Println("ERROR getCreate error: ", err)
			return
		}

		t.Execute(w, data)
	}
}
