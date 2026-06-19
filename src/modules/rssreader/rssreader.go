package rssreader

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func ParseRSS(file *os.File) ([]string, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	re := regexp.MustCompile(`<!\[CDATA\[(.*?)\]\]>`)
	matches := re.FindAllStringSubmatch(string(content), -1)

	var results []string

	for _, match := range matches {
		results = append(results, match[1])
	}

	return results, nil
}

func GetRSS(url string, filename string) (*os.File, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Error getting RSS content: %w", err)
	}

	defer res.Body.Close()

	buffer := make([]byte, 42)

	_, err = io.ReadFull(res.Body, buffer)

	if err != nil {
		return nil, err
	}

	signature := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><rss"

	if string(buffer) != signature {
		return nil, fmt.Errorf(
			"Invalid RSS signature, received: %s\n expected %s",
			buffer, signature)
	}

	file, err := os.Create(filename)

	if err != nil {
		file.Close()
		return nil, err
	}

	fullReader := io.MultiReader(bytes.NewReader(buffer), res.Body)

	_, err = io.Copy(file, fullReader)
	if err != nil {
		return nil, err
	}

	file.Seek(0, 0)

	return file, nil
}
