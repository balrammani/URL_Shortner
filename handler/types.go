package handler

import (
	"net/http"

	"github.com/balram/rest-url-shortner/controller"
)

type (
	// Handler ...
	Handler interface {
		URLShortner(w http.ResponseWriter, r *http.Request)
	}

	handler struct {
		controller controller.Controller
	}
)