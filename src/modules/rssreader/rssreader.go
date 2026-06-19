package rssreader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetRSS(url string, filename string) (*os.File, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Error getting RSS content: %w", err)
	}

	defer res.Body.Close()

	file, err := os.Create(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	_, err = io.Copy(file, res.Body)

	if err != nil {
		return nil, err
	}

	return file, nil
}
