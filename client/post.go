package client

import (
	"bytes"
	"net/http"
)

// PostRequest strategy
type PostRequest struct{}

// Do makes a generic POST request to the specified URL with the provided body and headers.
// It returns a Response struct and any error encountered.
func (p *PostRequest) Do(url string, body []byte, headers map[string]string, timeout int) (*Response, error) {
	reqBody := bytes.NewBuffer(body)
	req, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := createClient(timeout)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return handleResponse(resp)
}
