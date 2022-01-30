package controller

import (
	"fmt"
	"log"

	"github.com/balram/rest-url-shortner/repository"
	randomgenerator "github.com/balram/rest-url-shortner/utils/random_generator"
)

func New() (Controller, error) {
	random, err := randomgenerator.New()
	if err != nil {
		log.Fatalf("Handler New function error in randomgenerator object creation :%s", err.Error())
		return nil, err
	}
	repo, err := repository.New()
	if err != nil {
		log.Fatalf("Handler New function error in repository object creation :%s", err.Error())
		return nil, err
	}
	return &controller{
		random: random,
		repo:   repo,
	}, nil
}

func (h *controller) CreateShortLink(originalURL string) (string, error) {
	data, err := h.repo.Read()
	if err != nil {
		log.Fatalf("error in reading existing url from file :%s", err.Error())
		return "", err
	}

	if value, ok := data[originalURL]; ok {
		log.Printf("shortlink %q already available for the url %q", value, originalURL)
		return fmt.Sprintf("%s%s", REDIRECT_URL, value), nil
	}

	shortURL := h.random.GenerateRandomString()
	err = h.repo.Write(originalURL, shortURL)
	if err != nil {
		log.Fatalf("Handler New function error in repository object creation :%s", err.Error())
		return "", err
	}

	log.Printf("shortlink %q created for the url %q", shortURL, originalURL)
	return fmt.Sprintf("%s%s", REDIRECT_URL, shortURL), nil
}