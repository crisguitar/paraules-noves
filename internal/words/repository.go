package words

import (
	"github.com/crisguitar/paraules-noves/internal/infrastructure"
)

type Repository interface {
	Save(entry Entry) error
	GetAll() (entries []Entry, err error)
}

type PgRepository struct {
	Db infrastructure.DB
}

func (repo PgRepository) Save(entry Entry) error {
	query := "INSERT INTO words (word, meaning) values (:word, :meaning)"
	if _, err := repo.Db.NamedExec(query, &entry); err != nil {
		return err
	}
	return nil
}

func (repo PgRepository) GetAll() ([]Entry, error) {
	entries := []Entry{}
	query := "SELECT * FROM words"
	if err := repo.Db.Select(&entries, query); err != nil {
		return nil, err
	}

	return entries, nil
}

func NewRepository(config infrastructure.DbConfig) Repository {
	db, _ := infrastructure.CreateDB(config)
	return PgRepository{Db: db}
}
