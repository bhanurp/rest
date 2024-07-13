package rest

type Client struct {
	URL       string
	Headers   map[string]string
	Body      []byte
	Timeout   int
	Strategy  RequestMethod
	basicAuth *basicAuth
}

type Response struct {
	StatusCode int
	Body       []byte
	Error      error
}

type basicAuth struct {
	username string
	password string
}

// RequestMethod defines the interface for HTTP request strategies.
type RequestMethod interface {
	Do(url string, body []byte, headers map[string]string, timeout int) (*Response, error)
}
