package words

import (
	"encoding/json"
	"github.com/crisguitar/paraules-noves/internal/common"
	"io/ioutil"
	"net/http"
)

type CreateWordHandler struct {
	Repository Repository
}

func (h CreateWordHandler) Handle(w http.ResponseWriter, req *http.Request) (interface{}, error) {
	b, _ := ioutil.ReadAll(req.Body)
	newEntry := Entry{}
	err := json.Unmarshal(b, &newEntry)
	if err != nil {
		return nil, common.AppError{Message: "Wrong body", HttpCode: 400}
	}
	h.Repository.Save(newEntry)
	w.WriteHeader(201)
	return nil, nil
}

func NewCreateWordHandler(repo Repository) http.Handler {
	return common.RequestHandler{
		Handler: CreateWordHandler{
			Repository: repo,
		},
	}
}
