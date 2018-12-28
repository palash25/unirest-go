package unirest

type UnirestClient struct {
	timeout int
	userAgent string
	defaultHeaders map[string]string
}

func NewClient(opts ) (*UnirestClient, error) {

}

func (c *UnirestClient) func Get(url string, headers map[string]interface{}, body interface{}, auth[string]string) (*Response, error) {

}

func (c *UnirestClient) func Post(url string, headers map[string]interface{}, body interface{}, auth[string]string) (*Response, error) {

}

func (c *UnirestClient) func Put(url string, headers map[string]interface{}, body interface{}, auth[string]string) (*Response, error) {

}

func (c *UnirestClient) func Delete(url string, headers map[string]interface{}, body interface{}, auth[string]string) (*Response, error) {

}

func (c *UnirestClient) func Patch(url string, headers map[string]interface{}, body interface{}, auth[string]string) (*Response, error) {

}