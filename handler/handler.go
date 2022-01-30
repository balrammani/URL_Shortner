package handler

import (
	"log"
	"net/http"

	"github.com/balram/rest-url-shortner/controller"
)

func New() (Handler, error) {
	controller, err := controller.New()
	if err != nil {
		log.Fatalf("Handler New function error in controller object creation :%s", err.Error())
		return nil, err
	}

	return &handler{
		controller: controller,
	}, nil
}

func (h *handler) URLShortner(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("only support post request\n"))
		return
	}
	url := r.URL.Query().Get("url")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("need to send the url that must be shorten\n"))
		return
	}
	key, err := h.controller.CreateShortLink(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't create short link with provided key\n"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(key + "\n"))

}
