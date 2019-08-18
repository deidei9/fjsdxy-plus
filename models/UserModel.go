package models

type User struct {
	Id       int    `orm:"auto;pk;index" json:"id"`
	Phone    string `orm:"size(11)" json:"phone"`
	Password string `orm:"size(32)" json:"password"`
	NickName string `orm:"size(30);column(nickname)" json:"nickName"`
	Avator   string `orm:"size(255)" json:"avator"`
	State    uint8  `orm:default(1)`
	Model
}

func NewUser() *User {
	return &User{}
}
