package main

import (
	"net/http"
	"valentine-quote/controller"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.Handle("/", controller.GetIndex())
	http.Handle("/create", controller.CreateMessageHandler())

	http.ListenAndServe(":3000", nil)
}
