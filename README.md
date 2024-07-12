# rest
A simple wrapper around http client to make restful requests easier.

- An example usage for a post request
```go
client := DefaultClient()
client.URL = "http://example.com"
client.Headers = map[string]string{
	"ContentType": "application/json"
    "Authorization":"Bearer token"
}
client.Body = bytes.NewReader([]byte(`{"key": "value"}`))
client.Strategy = PostRequest{}
client.Send()
```
