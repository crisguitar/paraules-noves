package mocks

import (
	"database/sql"
	"github.com/stretchr/testify/mock"
)

type FakeDB struct {
	mock.Mock
}

func (db *FakeDB) NamedExec(query string, params interface{}) (sql.Result, error) {
	args := db.Called(query, params)
	return nil, args.Error(1)
}
