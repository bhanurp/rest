package rest

func (r *Response) IsResponseOK() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

func (r *Response) IsResponseError() bool {
	return r.StatusCode >= 400
}
