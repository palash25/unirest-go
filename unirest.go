package unirest

type UnirestClient struct {
	timeout        int
	userAgent      string
	defaultHeaders map[string]string
}

func NewClient() *UnirestClient {
	return &UnirestClient{
		timeout:   10,
		userAgent: "unirest-go/1.0",
	}
}

func (c *UnirestClient) Get(url string, headers map[string]interface{}, body interface{}, auth [string]string) (*Response, error) {
	r := NewRequest("GET", url, headers, body, auth)
	r.Do()
}

func (c *UnirestClient) Post(url string, headers map[string]interface{}, body interface{}, auth [string]string) (*Response, error) {
	r := NewRequest("GET", url, headers, body, auth)
	r.Do()
}

func (c *UnirestClient) Put(url string, headers map[string]interface{}, body interface{}, auth [string]string) (*Response, error) {
	r := NewRequest("GET", url, headers, body, auth)
	r.Do()
}

func (c *UnirestClient) Delete(url string, headers map[string]interface{}, body interface{}, auth [string]string) (*Response, error) {
	r := NewRequest("GET", url, headers, body, auth)
	r.Do()
}

func (c *UnirestClient) Patch(url string, headers map[string]interface{}, body interface{}, auth [string]string) (*Response, error) {
	r := NewRequest("GET", url, headers, body, auth)
	r.Do()
}

func (c *UnirestClient) SetTimeout(timeout int) *UnirestClient {
	c.timeout = timeout
	return c
}

func (c *UnirestClient) SetUserAgent(agent string) *UnirestClient {
	c.userAgent = agent
	return c
}

func (c *UnirestClient) SetDefaultHeader(key, value string) *UnirestClient {
	c.defaultHeaders[key] = value
	return c
}
