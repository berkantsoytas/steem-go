package jsonrpc2

import (
	"errors"
	"fmt"

	"github.com/monaco-io/request"
)

type JSONRPC2 struct {
	Version string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      string        `json:"id"`
}

type JSONRPC2Response struct {
	Version string         `json:"jsonrpc"`
	Result  interface{}    `json:"result,omitempty"`
	Error   *JSONRPC2Error `json:"error,omitempty"`
	ID      string         `json:"id"`
}

type JSONRPC2Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JSONRPC2Client struct {
	URL  string
	Data JSONRPC2
}

func (c *JSONRPC2Client) NewRequest(method string, params []interface{}) {
	c.Data = JSONRPC2{
		ID:      "1",
		Version: "2.0",
		Method:  method,
		Params:  params,
	}
}

func (e *JSONRPC2Error) Error() string {
	return e.Message
}

func (c *JSONRPC2Client) Send() (*JSONRPC2Response, error) {
	if c.Data.ID == "" {
		return nil, errors.New("request ID is empty")
	}

	resp, err := SendRequest(c.URL, c.Data)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	if resp.Error != nil {
		return nil, &JSONRPC2Error{
			Code:    resp.Error.Code,
			Message: resp.Error.Message,
			Data:    resp.Error.Data,
		}
	}

	return resp, nil
}

func SendRequest(url string, data JSONRPC2) (*JSONRPC2Response, error) {
	var result JSONRPC2Response

	resp := request.
		New().
		POST(url).
		AddHeader(map[string]string{
			"Content-Type": "application/json",
		}).
		AddJSON(data).
		Send().
		Scan(&result)

	if !resp.OK() {
		return nil, resp.Error()
	}

	return &result, nil
}

func NewClient(url string) *JSONRPC2Client {
	return &JSONRPC2Client{
		URL: url,
	}
}
