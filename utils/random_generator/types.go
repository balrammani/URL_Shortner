package randomgenerator

type (
	RandomService interface {
		GenerateRandomString() string
	}

	randomService struct {
		NumberOfDigits int
	}
)

const (
	shortURLLength = 10
)