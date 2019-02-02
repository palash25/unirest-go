package unirest

import (
	"net/http"
	"sync"
)

// UnirestClient defines the basic options for a client
type UnirestClient struct {
	Timeout        int
	UserAgent      string
	DefaultHeaders map[string]string
	InChan         chan *Request
	OutChan        chan *AsyncRequest
}

// NewClient is a constructor that initializes unirest client.
func NewClient() *UnirestClient {
	return &UnirestClient{
		Timeout:        10,
		UserAgent:      "unirest-go/1.0",
		DefaultHeaders: make(map[string]string),
		InChan:         make(chan *Request),
		OutChan:        make(chan *AsyncRequest),
	}
}

// Get makes get requests using the UnirestClient
func (uc *UnirestClient) Get(url string, headers map[string]interface{}, body interface{}, auth map[string]string) *Request {
	r := NewRequest("GET", url, headers, body, auth)
	return r.BuildRequest(uc)
}

// Post makes get requests using the UnirestClient
func (uc *UnirestClient) Post(url string, headers map[string]interface{}, body interface{}, auth map[string]string) *Request {
	r := NewRequest("POST", url, headers, body, auth)
	return r.BuildRequest(uc)
}

// Put makes get requests using the UnirestClient
func (uc *UnirestClient) Put(url string, headers map[string]interface{}, body interface{}, auth map[string]string) *Request {
	r := NewRequest("PUT", url, headers, body, auth)
	return r.BuildRequest(uc)
}

// Delete makes get requests using the UnirestClient
func (uc *UnirestClient) Delete(url string, headers map[string]interface{}, body interface{}, auth map[string]string) *Request {
	r := NewRequest("DELETE", url, headers, body, auth)
	return r.BuildRequest(uc)
}

// Patch makes get requests using the UnirestClient
func (uc *UnirestClient) Patch(url string, headers map[string]interface{}, body interface{}, auth map[string]string) *Request {
	r := NewRequest("PATCH", url, headers, body, auth)
	return r.BuildRequest(uc)
}

func (uc *UnirestClient) SetTimeout(timeout int) {
	uc.Timeout = timeout
}

func (uc *UnirestClient) SetUserAgent(agent string) {
	if agent != "" {
		uc.UserAgent = agent
	}
}

func (uc *UnirestClient) SetDefaultHeader(key, value string) {
	uc.DefaultHeaders[key] = value
}

func (uc *UnirestClient) ClearDefaultHeader() {
	uc.DefaultHeaders = make(map[string]string)
}

// Do performs synchronous requests
func (r *Request) Do() (*http.Response, error) {
	resp, err := r.HTTPClient.Do(r.HTTPRequest)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DoAsync performs requests asynchronously that it receives from the input
// channel and sends the response or errors to the output channel
func (uc *UnirestClient) DoAsync() {
	wg := sync.WaitGroup{}

	for req := range uc.InChan {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := req.HTTPClient.Do(req.HTTPRequest)
			//resp := NewResponse(res)
			asyncResp := &AsyncRequest{
				Resp: res,
				Err:  err,
			}
			uc.OutChan <- asyncResp
		}()
	}

	go func() {
		wg.Wait()
		close(uc.OutChan)
	}()
}
