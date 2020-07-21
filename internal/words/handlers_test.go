package words_test

import (
	"bytes"
	"encoding/json"
	"github.com/crisguitar/paraules-noves/internal/words"
	"github.com/crisguitar/paraules-noves/internal/words/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateWordHandler_Handle_Returns201(t *testing.T) {
	repo := new(mocks.FakeRepository)
	requestBody := map[string]string{"word": "some word", "meaning": "some meaning"}
	request, writer := prepareRequest(requestBody)
	repo.On("Save", mock.Anything).Return(nil)

	words.NewCreateWordHandler(repo).ServeHTTP(writer, request)

	code := writer.Code
	assert.Equal(t, 201, code)
}

func TestCreateWordHandler_Handle_ShouldSaveWord(t *testing.T) {
	repo := new(mocks.FakeRepository)
	word := "some word"
	meaning := "some meaning"
	requestBody := map[string]string{"word": word, "meaning": meaning}
	request, writer := prepareRequest(requestBody)
	repo.On("Save", mock.Anything).Return(nil)

	words.NewCreateWordHandler(repo).ServeHTTP(writer, request)

	repo.AssertCalled(t, "Save", words.Entry{Word: word, Meaning: meaning})
}

func TestCreateWordHandler_Handle_ShouldReturnErrorWithBadBody(t *testing.T) {
	repo := new(mocks.FakeRepository)
	requestBody := "this is wrong"
	request, writer := prepareRequest(requestBody)

	words.NewCreateWordHandler(repo).ServeHTTP(writer, request)

	response := make(map[string]string)
	body, _ := ioutil.ReadAll(writer.Body)
	json.Unmarshal(body, &response)

	code := writer.Code
	assert.Equal(t, 400, code)
	assert.Equal(t, "Wrong body", response["error"])
}

func TestCreateWordHandler_Handle_ShouldReturnErrorWhenInvalid(t *testing.T) {
	repo := new(mocks.FakeRepository)
	requestBody := map[string]string{}
	request, writer := prepareRequest(requestBody)

	words.NewCreateWordHandler(repo).ServeHTTP(writer, request)

	response := make(map[string]string)
	body, _ := ioutil.ReadAll(writer.Body)
	json.Unmarshal(body, &response)

	code := writer.Code
	assert.Equal(t, 400, code)
	assert.Equal(t, "Wrong body", response["error"])
}

func prepareRequest(requestBody interface{}) (*http.Request, *httptest.ResponseRecorder) {
	serialised, _ := json.Marshal(requestBody)
	request, _ := http.NewRequest("POST", "/words", bytes.NewReader(serialised))
	writer := httptest.NewRecorder()
	return request, writer
}
