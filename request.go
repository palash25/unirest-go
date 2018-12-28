package unirest

type Request struct {
	method  string
	url     string
	headers map[string]interface{}
	body    map[string]interface{}
}

func NewRequest(method string, url string, headers map[string]interface{}, body interface{}, auth [string]string) *Request {
	uniReq := Request{
		method:  method,
		url:     url,
		headers: headers,
		body:    body,
	}
	return uniReq
}
