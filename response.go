package unirest

import (
	"io/ioutil"
	"net/http"
)

// Response struct for the unirest protocol
type Response struct {
	Code    int
	Body    string
	RawBody string
	Headers map[string][]string
}

// NewResponse initializes and returns a unirest response struct or an error if any
func NewResponse(resp *http.Response) (*Response, error) {
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	unirestRespose := &Response{
		Code:    resp.StatusCode,
		Body:    string(responseData),
		RawBody: string(responseData),
		Headers: resp.Header,
	}
	return unirestRespose, nil
}
