package http

import (
	"fmt"
	"io"
	"net/http"
	"time"

	xerror "github.com/pkg/errors"
)

type requestBody struct {
	baseUrl     string
	url         string
	method      string
	contentType string
	Requests    *http.Request
	Response    *http.Response
}

func NewRequest(baseUrl string) *requestBody {
	return &requestBody{
		baseUrl:     baseUrl,
		url:         baseUrl,
		method:      "GET",
		contentType: "application/json",
	}
}

func (r *requestBody) SetUri(uri string) {
	r.url = fmt.Sprintf("%s%s", r.baseUrl, uri)
}

func (r *requestBody) GetURL() string {
	return r.url
}
func (r *requestBody) SetMethod(method string) {

	switch {
	case (method == "get" || method == "GET"):
		r.method = "GET"

	case (method == "post" || method == "POST"):
		r.method = "POST"
	default:
		r.method = "GET"
	}
}

// common get
func (r *requestBody) Get() error {
	client := NewHTTPClient()
	resp, err := client.Get(r.baseUrl)
	if err != nil {
		return xerror.Wrapf(err, "http request:%s failed. Cause:", r.baseUrl)
	}

	r.Response = resp
	return nil

}

// common post
// io.Reader :
// strings.NewReader("name=cjb")
// bytes.NewReader(data []byte)

func (r *requestBody) Post(body io.Reader) error {
	client := NewHTTPClient()
	resp, err := client.Post(r.url, r.contentType, body)
	if err != nil {
		return xerror.Wrapf(err, "http request:%s failed. Cause:", r.url)
	}

	r.Response = resp
	return nil
}

// common head
func (r *requestBody) Head() error {
	client := NewHTTPClient()
	resp, err := client.Head(r.url)
	if err != nil {
		return xerror.Wrapf(err, "http request:%s failed. Cause:", r.url)
	}

	r.Response = resp
	return nil
}

// a sp request
func (r *requestBody) Request(body io.Reader) error {
	req, err := http.NewRequest(r.method, r.url, body)
	if err != nil {
		return xerror.Wrapf(err, "prepare the http request(%s %s )failed. Cause:", r.url, r.method)
	}

	r.Requests = req
	return nil
}

// set the header
// Content-Type: application/json| application/x-www-form-urlencoded
//
func (r *requestBody) SetHeader(key, value string) {
	r.Requests.Header.Set(key, value)
}

func (r *requestBody) SetBasicAuth(name, password string) {
	r.Requests.SetBasicAuth(name, password)
}

func (r *requestBody) UserAgent() string {
	return r.Requests.UserAgent()
}

func (r *requestBody) Do() error {
	client := NewHTTPClient()
	resp, err := client.Do(r.Requests)
	if err != nil {
		return xerror.Wrapf(err, "failed to exec the http request for %s with:%v. Cause:", r.url, r.method)
	}

	r.Response = resp
	return nil
}

//

/*
	func (c *Client) Do(req *Request) (*Response, error)
	func (c *Client) Get(url string) (resp *Response, err error)
	func (c *Client) Head(url string) (resp *Response, err error)
	func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error)
	func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)
*/
func NewHTTPClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	return &http.Client{
		Timeout:   60 * time.Second,
		Transport: tr,
	}
}
