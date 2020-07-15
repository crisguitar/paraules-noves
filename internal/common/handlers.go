package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestHandler struct {
	Handler Handler
}

func (r RequestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := r.Handler.Handle(w, req)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if result == nil {
		return
	}

	body, err := json.Marshal(result)

	if err != nil {
		fmt.Println("Error marshaling JSON :(")
		return
	}

	w.Write(body)
}

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request) (interface{}, error)
}

type ApiError struct {
	Message string
}

func (e ApiError) Error() string {
	return e.Message
}
