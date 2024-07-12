package client

import (
	"io"
	"net/http"
	"time"
)

func createClient(timeout int) *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}

func handleResponse(resp *http.Response) (*Response, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &Response{
		StatusCode: resp.StatusCode,
		Body:       body,
		Error:      nil, // Error handling could be more nuanced
	}, nil
}
