package repository

type (
	// Repository ...
	Repository interface {
		Read() (map[string]string, error)
		Write(originalURL, shortURL string) error
	}

	repository struct {
	}
)
