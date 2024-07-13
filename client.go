package rest

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

// Send delegates the request handling to the strategy
func (c *Client) Send(url string, body []byte, headers map[string]string, timeout int) (*Response, error) {
	return c.Strategy.Do(url, body, headers, timeout)
}
