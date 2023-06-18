package client_services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	Difficulty = 4
	ServerHost = "http://quote_server:8081/quote"
)

type ServiceInterface interface {
	GetQuote() (string, error)
}

type Service struct{}

func (c *Service) GetQuote() (string, error) {
	resp, err := http.Get(ServerHost)
	if err != nil {
		return "", errors.New("failed to get challenge")
	}
	defer resp.Body.Close()

	challenge := resp.Header.Get("Challenge")
	if challenge == "" {
		return "", errors.New("failed to retrieve challenge")
	}

	solution := generateSolution(challenge)
	hashCash := fmt.Sprintf("%s:%s", challenge, solution)

	client := &http.Client{}
	req, err := http.NewRequest("GET", ServerHost, nil)
	if err != nil {
		return "", errors.New("failed to get quote")
	}

	req.Header.Set("HashCash", hashCash)

	resp, err = client.Do(req)
	if err != nil {
		return "", errors.New("failed to do request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("failed to read request")
	}

	return string(body), nil
}

func generateSolution(challenge string) string {
	for i := 0; ; i++ {
		solution := fmt.Sprintf("%s%d", challenge, i)
		hash := sha256.Sum256([]byte(solution))
		stringHash := hex.EncodeToString(hash[:])

		if strings.HasPrefix(stringHash, strings.Repeat("0", Difficulty)) {
			return solution
		}
	}
}
