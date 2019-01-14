package unirest

type Response struct {
	Code    int
	Body    string
	RawBody string
	Headers map[string][]string
}

/*
func NewResponse(code int, body, raw string, headers) (*Response, error) {
	return &Response{
		Code: code,
		Body: body,
		RawBody: raw,
		Headers: headers,
	}
}*/
