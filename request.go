package unirest

import (
	"encoding/base64"
	"net/http"
	"reflect"
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
		HTTPClient: &http.Client{
			Transport: nil,
		},
		HTTPRequest: &http.Request{
			Proto: "HTTP/1.0",
		},
	}
	return &uniReq
}

// BuildRequest performs the actual HTTP request after encoding the params and setting
// the required request headers.
func (r *Request) BuildRequest(uc *UnirestClient) *Request {
	// In case of GET request the params are encoded into
	// url otherwise we use form-encoded params or multipart
	// form-encoded based on whether the params contain a file
	// to be uploaded or not.
	if r.method == "GET" {
		r.rawEncode()
	} else {
		bodyMap, _ := r.body.(map[string]interface{})
		for key, value := range bodyMap {
			if key == "file" {
				delete(bodyMap, key)
				r.multiPartFormEncode(key, value.(string), bodyMap)
				break
			}
		}
		if r.HTTPRequest != nil {
			r.formEncode()
		}
	}

	r.HTTPRequest.Header.Set("user-agent", uc.UserAgent)
	r.HTTPRequest.Header.Set("accept-encoding", "gzip")
	// Setting headers received from the user to the
	// actual `Header` struct defined in the `net/http` package
	var vType string
	for k, v := range r.headers {
		vType = reflect.TypeOf(v).Kind().String()
		if vType != "string" {
			v = ToString(reflect.ValueOf(v), vType)
		}
		r.HTTPRequest.Header.Set(strings.ToLower(k), v.(string))
	}

	// Set basic auth header
	if r.auth != nil {
		data := []byte(r.auth["user"] + ":" + r.auth["password"])
		str := "Basic " + base64.StdEncoding.EncodeToString(data)
		r.HTTPRequest.Header.Set("Authorization", str)
	}

	return r
}
