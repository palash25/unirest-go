package unirest

import (
	"io/ioutil"
	"net/http"
)

// Response struct for the unirest protocol.
type Response struct {
	Code    int
	RawBody string
	Headers map[string][]string
}

// AsyncRequest is used for returning responses of asynchronous requests
// through channels.
type AsyncRequest struct {
	Resp *Response
	Err  error
}

// NewResponse initializes and returns a unirest response struct or an error if any
func NewResponse(resp *http.Response) *Response {
	responseData, _ := ioutil.ReadAll(resp.Body)

	unirestRespose := &Response{
		Code:    resp.StatusCode,
		RawBody: string(responseData),
		Headers: resp.Header,
	}
	return unirestRespose
}
