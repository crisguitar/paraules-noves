package words

import "testing"

func TestEntry_IsValid(t *testing.T) {
	type fields struct {
		Word    string
		Meaning string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{ "when word is empty", fields{Word: "", Meaning: "some meaning"}, false},
		{ "when meaning is empty", fields{Word: "word", Meaning: ""}, false},
		{ "when valid", fields{Word: "word", Meaning: "meaning"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Entry{
				Word:    tt.fields.Word,
				Meaning: tt.fields.Meaning,
			}
			if got := e.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
