package common_test

import (
	"encoding/json"
	"github.com/crisguitar/paraules-noves/internal/common"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakeSuccessHandler struct {
	Result interface{}
}
type FakeErrorHandler struct {
	Message string
}

func (h FakeSuccessHandler) Handle(_ http.ResponseWriter, _ *http.Request) (interface{}, error) {
	return h.Result, nil
}

func (h FakeErrorHandler) Handle(_ http.ResponseWriter, _ *http.Request) (interface{}, error) {
	return nil, common.ApiError{
		Message: h.Message,
	}
}

func TestRequestHandler_ServeHTTP_setsJsonHeader(t *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/does-not-matter", nil)
	handler := common.RequestHandler{
		Handler: FakeSuccessHandler{},
	}

	handler.ServeHTTP(writer, request)

	assert.Equal(t, writer.Header().Get("Content-Type"), "application/json")
	assert.Equal(t, writer.Code, 200)
}

func TestRequestHandler_ServeHTTP_returnsResponse(t *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/does-not-matter", nil)
	expectedResponse := map[string]string{"key": "value"}
	handler := common.RequestHandler{
		Handler: FakeSuccessHandler{
			Result: expectedResponse,
		},
	}

	handler.ServeHTTP(writer, request)

	response := make(map[string]string)
	bytes, _ := ioutil.ReadAll(writer.Body)
	json.Unmarshal(bytes, &response)
	assert.Equal(t, response, expectedResponse)
}

func TestRequestHandler_ServeHTTP_whenHandlerFailsReturnError(t *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/does-not-matter", nil)
	handler := common.RequestHandler{
		Handler: FakeErrorHandler{
			Message: "sad error",
		},
	}

	handler.ServeHTTP(writer, request)

	assert.Equal(t, 500, writer.Code)
}
