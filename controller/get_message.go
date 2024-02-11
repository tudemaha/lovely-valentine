package controller

import (
	"log"
	"net/http"
	"strings"
	"text/template"
	"valentine-quote/model"
)

func GetMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paths := strings.Split(r.URL.Path, "/")
		id := paths[1]
		var message model.Message

		message, available := model.CheckAvailableID(id)
		if !available || len(paths) != 2 {
			http.Redirect(w, r, "/create", http.StatusSeeOther)
		}

		t, err := template.ParseFiles("./view/finish.html")
		if err != nil {
			log.Println("ERROR getCreate error: ", err)
			return
		}

		t.Execute(w, message)
	}
}
