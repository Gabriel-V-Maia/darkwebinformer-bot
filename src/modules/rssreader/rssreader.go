package rssreader

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Category    string `xml:"category"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
}

func ParseRSS(file *os.File) ([]Item, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var rss RSS
	err = xml.Unmarshal(content, &rss)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling xml: %w", err)
	}

	return rss.Channel.Items, nil
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
