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
		writeError(w, err)
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

func writeError(w http.ResponseWriter, err error) {
	var errorInBytes []byte
	if appError, isAppError := err.(AppError); isAppError {
		w.WriteHeader(appError.HttpCode)
		errorInBytes, _ = json.Marshal(appError)
	} else {
		w.WriteHeader(500)
		errorInBytes, _ = json.Marshal(AppError{
			HttpCode: 500,
			Message:  err.Error(),
		})
	}

	w.Write(errorInBytes)
}
