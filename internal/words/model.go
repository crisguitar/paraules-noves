package words

type Entry struct {
	Id      int    `db:"id" json:"id"`
	Word    string `db:"word" json:"word"`
	Meaning string `db:"meaning" json:"meaning"`
}

func (e Entry) IsValid() bool {
	return e.Meaning != "" && e.Word != ""
}
