package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./view/index.html")
		if err != nil {
			fmt.Println("error: ", err)
			return
		}

		data := struct {
			QuotesNum []uint8
			Quotes    []string
		}{
			QuotesNum: []uint8{0, 1, 2, 3, 4},
			Quotes:    []string{"hai", "ini", "test", "dinisi", "testlagi"},
		}

		t.Execute(w, data)
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./view/create.html")
		if err != nil {
			fmt.Println("error: ", err)
			return
		}

		data := struct {
			QuotesNum []uint8
		}{
			QuotesNum: []uint8{1, 2, 3, 4, 5},
		}

		t.Execute(w, data)
	})

	http.HandleFunc("/finish", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./view/finish.html")
		if err != nil {
			fmt.Println("error: ", err)
			return
		}

		data := struct {
			Status bool
		}{
			Status: true,
		}

		t.Execute(w, data)
	})

	http.HandleFunc("/input", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./view/password.html")
		if err != nil {
			fmt.Println("error: ", err)
			return
		}

		t.Execute(w, nil)
	})

	http.ListenAndServe(":3000", nil)
}
