package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type QuoteInterface interface {
	GetRandomQuote() (*Quote, error)
}

type QuoteService struct {
	Link   string
	Method string
	Format string
	Lang   string
}

type Quote struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
}

func NewQuoteService(link, method, format, lang string) *QuoteService {
	return &QuoteService{
		Link:   link,
		Method: method,
		Format: format,
		Lang:   lang,
	}
}

func (c *QuoteService) GetRandomQuote() (*Quote, error) {
	request := fmt.Sprintf("%s?method=%s&format=%s&lang=%s", c.Link, c.Method, c.Format, c.Lang)
	response, err := http.Get(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var quote Quote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil, err
	}

	return &quote, nil
}
