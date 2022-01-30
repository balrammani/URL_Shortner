package controller

import (
	"github.com/balram/rest-url-shortner/repository"
	randomgenerator "github.com/balram/rest-url-shortner/utils/random_generator"
)

type (
	// Controller ...
	Controller interface {
		CreateShortLink(url string) (string, error)
		Redirect(url string) (string, error)
	}

	controller struct {
		repo   repository.Repository
		random randomgenerator.RandomService
	}
)

const (
	REDIRECT_URL = "localhost:8080/"
)


