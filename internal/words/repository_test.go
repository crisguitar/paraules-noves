package words_test

import (
	"github.com/crisguitar/paraules-noves/internal/common"
	"github.com/crisguitar/paraules-noves/internal/words"
	"github.com/crisguitar/paraules-noves/internal/words/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestRepository_SaveWord(t *testing.T) {
	var fakeDb = mocks.FakeDB{}
	fakeDb.On("NamedExec", mock.Anything, mock.Anything).Return(nil, nil)
	repo := words.PgRepository{
		Db: &fakeDb,
	}
	word := "some word"
	meaning := "some meaning"
	entry := words.Entry{
		Word:    word,
		Meaning: meaning,
	}

	err := repo.Save(entry)

	expectedQuery := "INSERT INTO words (word, meaning) values (:word, :meaning)"
	fakeDb.AssertCalled(t, "NamedExec", expectedQuery, &entry)
	assert.Nil(t, err)
}

func TestRepository_ReturnsErrorIfDBFails(t *testing.T) {
	var fakeDb = mocks.FakeDB{}
	expectedError := common.AppError{}
	fakeDb.On("NamedExec", mock.Anything, mock.Anything).Return(nil, expectedError)
	repo := words.PgRepository{
		Db: &fakeDb,
	}
	entry := words.Entry{
		Word:    "some word",
		Meaning: "some meaning",
	}

	err := repo.Save(entry)

	assert.Equal(t, expectedError, err)
}
