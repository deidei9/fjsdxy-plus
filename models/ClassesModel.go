package models

type Classes struct {
	Id        int       `orm:"auto;pk;index"`
	ClassName string    `orm:"size(30)"`
	College   string    `orm:"size(30)"`
	Major     string    `orm:"size(30)"`
	Course    []*Course `orm:"reverse(many);on_delete(do_nothing)"`
	Model
}

func NewClasses() *Classes {
	return &Classes{}
}
