package main

import (
	"log"
	"net/http"

	"github.com/balram/rest-url-shortner/handler"
	"github.com/gorilla/mux"
)

func main(){
	log.Println("Starting rest service - url shortner")

	handlers, err := handler.New()
	if err != nil {
		log.Fatalf("Handler New function error in main method :%s", err.Error())
	}

	r := mux.NewRouter()
	r.HandleFunc("/create_short_link", handlers.URLShortner)
	r.HandleFunc("/{shortlink}", handlers.Redirect)
	log.Fatal( http.ListenAndServe(":8080", r))
}