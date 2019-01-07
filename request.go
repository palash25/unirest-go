package unirest

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// Request defines the contents of a unirest-request
type Request struct {
	method      string
	url         string
	headers     map[string]interface{}
	body        interface{}
	auth        map[string]string
	HTTPRequest *http.Request
	HTTPClient  *http.Client
}

// NewRequest is a constructor that initializes the Request struct.
func NewRequest(method, url string, headers map[string]interface{}, body interface{}, auth map[string]string) *Request {
	uniReq := Request{
		method:  method,
		url:     url,
		headers: headers,
		body:    body,
		auth:    auth,
	}
	return &uniReq
}

// Do performs the actual HTTP request after encoding the params and setting
// the required request headers.
func (r *Request) Do() {
	// In case of GET request the params are encoded into
	// url otherwise we use form-encoded params or multipart
	// form-encoded based on whether the params contain a file
	// to be uploaded or not.
	if r.method == "GET" {
		r.rawEncode()
		r.HTTPRequest, _ = http.NewRequest("GET", r.url, nil)
	} else {
		bodyMap, _ := r.body.(map[string]interface{})
		for key, value := range bodyMap {
			if key == "file" {
				delete(bodyMap, key)
				r.multiPartFormEncode(key, value.(string), bodyMap)
			}
		}
		if r.HTTPRequest != nil {
			r.formEncode()
		}
	}

	// Setting headers received from the user to the
	// actual `Header` struct defined in the `net/http` package
	for k, v := range r.headers {
		r.HTTPRequest.Header.Set(strings.ToLower(k), v.(string))
	}

	if r.auth != nil {
		data := []byte(r.auth["user"] + ":" + r.auth["password"])
		str := base64.StdEncoding.EncodeToString(data)
		r.HTTPRequest.Header.Set("Authorization", "Basic"+str)
	}
}
