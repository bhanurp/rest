package rest

import "net/url"

// NewClient creates a new Client struct with the provided URL and headers.
// the default timeout is 30 seconds
// the default method is GET
// the default Body is an empty byte slice
func NewClient(url string, headers map[string]string) *Client {
	return &Client{
		URL:      url,
		Headers:  headers,
		Body:     []byte{},
		Timeout:  10,
		Strategy: &GetRequest{},
	}
}

// DefaultClient creates a new Client struct with default values.
// URL: empty string
// headers: nil
// timeout: 10 seconds
// Body: empty byte slice
// Strategy: nil
func DefaultClient() *Client {
	return &Client{
		URL:      "",
		Headers:  nil,
		Body:     []byte{},
		Timeout:  10,
		Strategy: nil,
	}
}

// CreateClientWithTimeout creates a new http.Client with the provided timeout
func CreateClientWithTimeout(timeout int) *Client {
	client := DefaultClient()
	client.Timeout = timeout
	return client
}

// SetStrategy allows setting the current request strategy
func (c *Client) SetStrategy(s RequestMethod) {
	c.Strategy = s
}

// SetTimeout sets a custom timeout for the HTTP client
func (c *Client) SetTimeout(timeout int) {
	c.Timeout = timeout
}

// SetBody sets a custom body for the request
func (c *Client) SetBody(body []byte) {
	c.Body = body
}

// Send delegates the request handling to the strategy
func (c *Client) Send() (*Response, error) {
	return c.Strategy.Do(c.URL, c.Body, c.Headers, c.Timeout)
}

// AddQueryParams adds query parameters to the given URL and returns the modified URL
func (c *Client) AddQueryParams(params map[string]string) (string, error) {
	u, err := url.Parse(c.URL)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// AddBasicAuth adds basic authentication headers to the request
func AddBasicAuth(username, password string, c *Client) {
	c.basicAuth = &basicAuth{username: username, password: password}
}
