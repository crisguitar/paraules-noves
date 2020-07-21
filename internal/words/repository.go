package words

import (
	"github.com/crisguitar/paraules-noves/internal/common"
	"github.com/crisguitar/paraules-noves/internal/words/infrastructure"
)

type Repository interface {
	Save(entry Entry) error
}

type PgRepository struct {
	Db common.DB
}

func (repo PgRepository) Save(entry Entry) error {
	query := "INSERT INTO words (word, meaning) values (:word, :meaning)"
	if _, err := repo.Db.NamedExec(query, &entry); err != nil {
		return err
	}
	return nil
}

func NewRepository(config infrastructure.DbConfig) Repository {
	db, _ := infrastructure.CreateDB(config)
	return PgRepository{Db: db}
}
