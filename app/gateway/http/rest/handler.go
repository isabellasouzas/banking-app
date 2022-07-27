package rest

import (
	"encoding/json"
	"net/http"
)

type Handler func(*http.Request) Response

func (handler Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	ch := make(chan Response)
	go func() { ch <- handler(r) }()
	res := <-ch

	if res.Headers != nil {
		header := rw.Header()
		for k, v := range res.Headers {
			header.Add(k, v)
		}
	}

	if res.HTTPStatus != 0 {
		rw.WriteHeader(res.HTTPStatus)
	}

	if res.Error != nil || res.Payload != nil {
		payload, err := json.Marshal(res)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = rw.Write(payload)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
