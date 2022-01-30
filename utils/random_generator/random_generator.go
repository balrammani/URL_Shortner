package randomgenerator

import (
	"math/rand"
	"time"
)

func New() (RandomService, error) {
	return &randomService{NumberOfDigits: shortURLLength}, nil
}

func (r *randomService) GenerateRandomString() string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, r.NumberOfDigits)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}