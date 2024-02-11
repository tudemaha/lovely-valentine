package controller

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/tudemaha/lovely-valentine/model"
)

func GetMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		paths := strings.Split(r.URL.Path, "/")
		id := paths[1]
		if id != "" {
			http.Redirect(w, r, "/password/"+id, http.StatusSeeOther)
			return
		}

		data := struct {
			id       string
			password string
		}{
			id:       r.PostForm.Get("id"),
			password: r.PostForm.Get("password"),
		}

		if data.id == "" || data.password == "" {
			http.Redirect(w, r, "/create", http.StatusSeeOther)
			return
		}

		message, available := model.CheckAvailableID(data.id)
		if !available {
			http.Redirect(w, r, "/create", http.StatusSeeOther)
			return
		}

		if message.HashedMagic != data.password {
			http.Redirect(w, r, "/password/"+data.id, http.StatusSeeOther)
			return
		}

		t, err := template.ParseFiles("./view/index.html")
		if err != nil {
			log.Println("ERROR getCreate error: ", err)
			return
		}

		t.Execute(w, message)
	}
}
