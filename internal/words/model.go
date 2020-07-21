package words

type Entry struct {
	Word    string `db:"word"`
	Meaning string `db:"meaning"`
}
