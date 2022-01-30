package repository

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	FilePath = "config/config.txt"
)

func init() {
	_, err := os.Stat(FilePath)
	if os.IsNotExist(err) {
		_, e := os.Create(FilePath)
		if e != nil {
			log.Fatal(e)
		}
	}
}

// New constructs a new repository
func New() (Repository, error) {
	return &repository{
	}, nil
}

func (r *repository) Read() (map[string]string, error) {
	m := make(map[string]string)
	file, err := os.Open(FilePath)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), ",")
		if len(tmp) > 1 {
			m[tmp[0]] = tmp[1]
		}
	}
	return m, nil
}

func (r *repository) Write(originalURL, shortURL string) error {
	file, err := os.OpenFile(FilePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("\n%s,%s", originalURL, shortURL))
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return err
	}
	return nil
}
