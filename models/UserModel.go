package models

import "errors"

type User struct {
	Id       int     `orm:"auto;pk;index" json:"id"`
	Phone    string  `orm:"null;size(11)" json:"phone"`
	Password string  `orm:"size(32)" json:"password"`
	NickName string  `orm:"size(30);column(nickname)" json:"nickName"`
	Avator   string  `orm:"size(255)" json:"avator"`
	State    float64 `orm:"digits(1);decimals(0);default(1)" json:"state"`
	Model
}

func NewUser() *User {
	return &User{}
}

//获取用户信息
func (this *User) GetInfo() (err error) {
	if err = DB().Read(this); err != nil {
		return errors.New("用户不存在")
	}
	return nil
}
