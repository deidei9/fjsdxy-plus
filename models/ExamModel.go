package models

type Exam struct {
	Id      int     `orm:"auto;pk;index" json:"id"`
	Term    string  `orm:"size(30)" json:"term"`
	Name    string  `orm:"size(30)" json:"name"`
	Type    string  `orm:"size(30)" json:"type"`
	Score   string  `orm:"size(5)" json:"score"`
	Credits float32 `json:"credits"`
	User    *User   `orm:"rel(fk)"`
	Model
}

func NewExam() *Exam {
	return &Exam{}
}
