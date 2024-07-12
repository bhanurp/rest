package client

type Client struct {
	URL      string
	Headers  map[string]string
	Body     []byte
	Timeout  int
	Strategy RequestMethod
}

type Response struct {
	StatusCode int
	Body       []byte
	Error      error
}

// RequestMethod defines the interface for HTTP request strategies.
type RequestMethod interface {
	Do(url string, body []byte, headers map[string]string, timeout int) (*Response, error)
}
