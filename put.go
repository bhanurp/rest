package rest

import (
	"bytes"
	"net/http"
)

// PutRequest strategy
type PutRequest struct{}

func (p *PutRequest) Do(url string, body []byte, headers map[string]string, timeout int) (*Response, error) {
	reqBody := bytes.NewBuffer(body)
	req, err := http.NewRequest(http.MethodPut, url, reqBody)
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
