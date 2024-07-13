package rest

import (
	"net/http"
)

type DeleteRequest struct{}

func (d *DeleteRequest) Do(url string, headers map[string]string, timeout int) (*Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
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
