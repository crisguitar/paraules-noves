package words

type Repository interface {
	Save(entry Entry) error
}
