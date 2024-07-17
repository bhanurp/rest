package rest

import (
	"net/http"
)

// HeadRequest strategy
type HeadRequest struct{}

// Do makes a generic HEAD request to the specified URL with the provided headers.
// It returns the response body as a byte slice and any error encountered.
func (h *HeadRequest) Do(url string, body []byte, headers map[string]string, timeout int) (*Response, error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}

	// Add headers to the request
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
