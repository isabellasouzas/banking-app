package rest

import (
	"net/http"
)

type Response struct {
	HTTPStatus int               `json:"status"`
	Error      error             `json:"error"`
	Payload    interface{}       `json:"payload"`
	Headers    map[string]string `json:"headers"`
}

type Error struct {
	Type   string `json:"type"`
	Tittle string `json:"tittle"`
}

type Header struct {
	Key   string
	Value string
}

func (e Error) Error() string {
	return e.Type
}

var (
	ErrAccountNotFound = Error{Type: "error:account_not_found", Tittle: "Invalid Account ID"}
	ErrInvalidParams   = Error{Type: "error:invalid_params", Tittle: "Invalid Parameters"}
)

type response struct {
	HTTPStatus int         `json:"http_status,omitempty"`
	Error      error       `json:"error,omitempty"`
	Content    interface{} `json:"content,omitempty"`
	headers    map[string]string
	err        error
}

func (r response) AddHeaders(key, value string) response {
	if r.headers == nil {
		r.headers = map[string]string{}
	}
	r.headers[key] = value
	return r
}

func Created(payload interface{}, headers ...Header) Response {
	return Response{
		HTTPStatus: http.StatusCreated,
		Payload:    payload,
		Headers:    nil,
	}
}

func OK(payload interface{}, headers ...Header) Response {
	return Response{
		HTTPStatus: http.StatusOK,
		Payload:    payload,
		Headers:    nil,
	}
}
