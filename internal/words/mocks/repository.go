package mocks

import (
	"github.com/crisguitar/paraules-noves/internal/words"
	"github.com/stretchr/testify/mock"
)

type FakeRepository struct {
	mock.Mock
}

func (r *FakeRepository) Save(entry words.Entry) error {
	args := r.Called(entry)
	return args.Error(0)
}

func (r *FakeRepository) GetAll() ([]words.Entry, error) {
	args := r.Called()
	return nil, args.Error(1)
}
