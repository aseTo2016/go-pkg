package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Cli struct {
	nativeClient *http.Client
}

type Builder struct {
	method string
	url    string
	body   interface{}

	statusCheck       func(resp *http.Response) error
	parseResponseBody func(resp *http.Response, result interface{}) error

	result interface{} //必须是指针
}

// StatusCheck sets response status check func
func (b *Builder) StatusCheck(statusCheck func(resp *http.Response) error) *Builder {
	b.statusCheck = statusCheck
	return b
}

// ParseResponseBody sets parse response body func
func (b *Builder) ParseResponseBody(parseResponseBody func(resp *http.Response, result interface{}) error) *Builder {
	b.parseResponseBody = parseResponseBody
	return b
}

func (b *Builder) Result(result interface{}) *Builder {
	b.result = result
	return b
}

// NewBuilder return a builder
func NewBuilder(url string) *Builder {
	return &Builder{
		url:               url,
		statusCheck:       defaultStatusCheck,
		parseResponseBody: defaultParseResponseBody,
	}
}

func defaultStatusCheck(resp *http.Response) error {
	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return fmt.Errorf("status code: %d, response body: %s", resp.StatusCode, string(data))
}

func defaultParseResponseBody(resp *http.Response, result interface{}) error {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, result)
}

func (c *Cli) Do(builder *Builder) (err error) {
	data, ok := builder.body.([]byte)
	if !ok {
		data, err = json.Marshal(builder.body)
		if err != nil {
			return err
		}
	}
	req, err := http.NewRequest(builder.method, builder.url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	resp, err := c.nativeClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		errClose := resp.Body.Close()
		if errClose != nil {
			if err == nil {
				err = errClose
				return
			}
			err = fmt.Errorf("%s, close error: %s", err.Error(), errClose.Error())
		}
	}()

	err = builder.statusCheck(resp)
	if err != nil {
		return err
	}

	if builder.result == nil {
		return
	}
	return builder.parseResponseBody(resp, builder.result)
}

func (c *Cli) Get(builder *Builder) error {
	builder.method = http.MethodGet
	return c.Do(builder)
}

func (c *Cli) Put(builder *Builder) error {
	builder.method = http.MethodPut
	return c.Do(builder)
}

// Post ...
func (c *Cli) Post(builder *Builder) error {
	builder.method = http.MethodPost
	return c.Do(builder)
}

// Delete ...
func (c *Cli) Delete(builder *Builder) error {
	builder.method = http.MethodDelete
	return c.Do(builder)
}

// Patch ...
func (c *Cli) Patch(builder *Builder) error {
	builder.method = http.MethodPatch
	return c.Do(builder)
}

// Option ...
func (c *Cli) Option(builder *Builder) error {
	builder.method = http.MethodOptions
	return c.Do(builder)
}
