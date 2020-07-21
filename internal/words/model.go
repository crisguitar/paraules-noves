package words

type Entry struct {
	Word    string `db:"word"`
	Meaning string `db:"meaning"`
}

func (e Entry) IsValid() bool {
	return e.Meaning != "" && e.Word != ""
}
