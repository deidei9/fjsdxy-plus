package models

type Wechat struct {
	Id         int    `orm:"auto;pk;index" json:"id"`
	Openid     string `orm:"size(255)" json:"openid"`
	SessionKey string `orm:"size(255)" json:"session_key"`
	NickName   string `orm:"size(30)" json:"nick_name"`
	AvatarUrl  string `orm:""size(255)" json:"avator_url"`
	Gender     string `orm:"size(3)" json:"gender"`
	User       *User  `orm:"rel(one)" json:"user_info"`
	Model
}

func NewWechat() *Wechat {
	return &Wechat{}
}
