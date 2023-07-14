package httpClient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Get[T any](url string, body *T) (*T, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	// retryAfter := 300
	// req.Header.Set("Retry-After", fmt.Sprintf("%d", retryAfter))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error in calling api %v", resp.StatusCode)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			log.Println("404")
		}
		return nil, fmt.Errorf("error in calling api with status %v", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(body); err != nil {
		return nil, fmt.Errorf("error in formatting response")
	}
	return body, nil
}
