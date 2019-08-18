package models

type Student struct {
	Id          int      `orm:"auto;pk;index"`
	StudentId   string   `orm:"size(30)"`
	Password    string   `orm:"size(30)"`
	StudentName string   `orm:"size(30)"`
	User        *User    `orm:"rel(one)" json:"user_info"`
	Classes     *Classes `orm:"rel(one)"`
	Model
}

func NewStudent() *Student {
	return &Student{}
}
