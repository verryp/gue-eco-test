package common

import (
	"net/http"
	"sync"
)

var oneRsp sync.Once
var rsp *Response

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type Request struct {
	Request *http.Request
	Config  *Config
}

func (res *Response) SetMessage(msg string) *Response {
	res.Message = msg
	return res
}

func (res *Response) SetStatus(s string) *Response {
	res.Status = s
	return res
}

func (res *Response) SetData(data interface{}) *Response {
	res.Data = data
	return res
}

func (res *Response) SetErrors(errs interface{}) *Response {
	res.Errors = errs
	return res
}

// NewResponse initialize response
func NewResponse() *Response {
	oneRsp.Do(func() {
		rsp = &Response{}
	})

	// clone response
	x := *rsp

	return &x
}
