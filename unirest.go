package unirest

// UnirestClient defines the basic options for a client
type UnirestClient struct {
	timeout        int
	userAgent      string
	defaultHeaders map[string]string
}

// NewClient is a constructor that initializes unirest client.
func NewClient() *UnirestClient {
	return &UnirestClient{
		timeout:        10,
		userAgent:      "unirest-go/1.0",
		defaultHeaders: make(map[string]string),
	}
}

// Get makes get requests using the UnirestClient
func (c *UnirestClient) Get(url string, headers map[string]interface{}, body interface{}, auth map[string]string) (*Response, error) {
	r := NewRequest("GET", url, headers, body, auth)
	r.MakeRequest()
	return nil, nil
}

// Get makes get requests using the UnirestClient
func (c *UnirestClient) Post(url string, headers map[string]interface{}, body interface{}, auth map[string]string) (*Response, error) {
	r := NewRequest("POST", url, headers, body, auth)
	r.MakeRequest()
	return nil, nil
}

// Get makes get requests using the UnirestClient
func (c *UnirestClient) Put(url string, headers map[string]interface{}, body interface{}, auth map[string]string) (*Response, error) {
	r := NewRequest("PUT", url, headers, body, auth)
	r.MakeRequest()
	return nil, nil
}

// Get makes get requests using the UnirestClient
func (c *UnirestClient) Delete(url string, headers map[string]interface{}, body interface{}, auth map[string]string) (*Response, error) {
	r := NewRequest("DELETE", url, headers, body, auth)
	r.MakeRequest()
	return nil, nil
}

// Get makes get requests using the UnirestClient
func (c *UnirestClient) Patch(url string, headers map[string]interface{}, body interface{}, auth map[string]string) (*Response, error) {
	r := NewRequest("PATCH", url, headers, body, auth)
	r.MakeRequest()
	return nil, nil
}

func (c *UnirestClient) SetTimeout(timeout int) {
	c.timeout = timeout
}

func (c *UnirestClient) SetUserAgent(agent string) {
	c.userAgent = agent
}

func (c *UnirestClient) SetDefaultHeader(key, value string) {
	c.defaultHeaders[key] = value
}
