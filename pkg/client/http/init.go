package httpClient

import (
	"net/http"
	"time"
)

type Config struct {
	Timeout time.Duration
}

var client *http.Client

func NewClient(config *Config) *http.Client {
	client := &http.Client{
		Timeout: config.Timeout,
	}
	return client
}

func init() {
	client = NewClient(&Config{
		30000000000,
	})
}
