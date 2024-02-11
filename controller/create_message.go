package controller

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"valentine-quote/model"
	"valentine-quote/utils"
)

func CreateMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCreate(&w)
		}

		if r.Method == "POST" {
			postCreate(&w, r)
		}
	}
}

func getCreate(w *http.ResponseWriter) {
	t, err := template.ParseFiles("./view/create.html")
	if err != nil {
		log.Println("ERROR getCreate error: ", err)
		return
	}

	data := struct {
		QuotesNum []uint8
	}{
		QuotesNum: []uint8{1, 2, 3, 4, 5},
	}

	t.Execute(*w, data)
}

func postCreate(w *http.ResponseWriter, r *http.Request) {
	var messages model.NewMessage

	if err := r.ParseForm(); err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}

	messages.Magic = r.PostForm.Get("magic")
	for i := 1; i <= 5; i++ {
		if r.PostForm.Get("quote"+fmt.Sprint(i)) != "" {
			messages.Quotes = append(messages.Quotes, r.PostForm.Get("quote"+fmt.Sprint(i)))
		}
	}
	messages.HashedMagic = utils.CreateMD5Hash(messages.Magic)

	status := true
	insertedID, err := model.CreateMessage(messages)
	if err != nil {
		log.Printf("ERROR CreateMessage fatal error: %v", err)
		status = false
	}

	urlScheme := "http://"
	if r.TLS != nil {
		urlScheme = "https://"
	}

	data := struct {
		Status   bool
		Link     string
		Password string
	}{
		Status:   status,
		Link:     urlScheme + r.Host + "/" + insertedID,
		Password: messages.HashedMagic,
	}

	t, errTemp := template.ParseFiles("./view/finish.html")
	if errTemp != nil {
		log.Println("ERROR getCreate error: ", err)
		return
	}

	t.Execute(*w, data)
}
