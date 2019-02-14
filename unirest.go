package unirest

import (
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

// Do performs the requests
func (r *Request) Do() (*Response, error) {
	res, err := r.httpClient.Do(r.httpRequest)
	if err != nil {
		return nil, err
	}
	resp := NewResponse(res)
	return resp, nil
}

// DoAsync performs requests asynchronously that it receives from the input
// channel and sends the response or errors to the output channel
func (uc *UnirestClient) DoAsync() {
	wg := sync.WaitGroup{}

	for req := range uc.InChan {
		wg.Add(1)
		go func(r *Request) {
			defer wg.Done()
			resp, err := r.Do()
			asyncResp := &AsyncRequest{
				Resp: resp,
				Err:  err,
			}
			uc.OutChan <- asyncResp
		}(req)
	}

	go func() {
		wg.Wait()
		close(uc.OutChan)
	}()
}
