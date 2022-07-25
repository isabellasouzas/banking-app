package rest

import "net/http"

type Response struct {
	Status  int               `json:"status"`
	Error   error             `json:"error"`
	Payload interface{}       `json:"payload"`
	Headers map[string]string `json:"headers"`
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

func Created(payload interface{}, headers ...Header) Response {
	return Response{
		Status:  http.StatusCreated,
		Payload: payload,
		Headers: nil,
	}
}

func OK(payload interface{}, headers ...Header) Response {
	return Response{
		Status:  http.StatusOK,
		Payload: payload,
		Headers: nil,
	}
}

func HandleError(err error) {

}
