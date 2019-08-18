package models

type Course struct {
	Id       int      `orm:"auto;pk;index"`
	Term     string   `orm:"size(30)"`
	Weekly   int      `orm:"unique"`
	Name     string   `orm:"size(255)"`
	Position string   `orm:"size(10)"`
	Class    string   `orm:"size(30)"`
	Teacher  string   `orm:"size(30)"`
	Classes  *Classes `orm:"rel(fk);on_delete(do_nothing)"`
	Model
}

func NewCourse() *Course {
	return &Course{}
}
