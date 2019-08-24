package models

import (
	"errors"
	"github.com/pig0224/fjsdxy/cas/course"
)

type Course struct {
	Id       int      `orm:"auto;pk;index"`
	Term     string   `orm:"size(30)" json:"-"`
	Weekly   int      `orm:"unique"`
	Name     string   `orm:"size(255)"`
	Position string   `orm:"size(10)"`
	Class    string   `orm:"size(30)"`
	Teacher  string   `orm:"size(30)"`
	Classes  *Classes `orm:"rel(fk);on_delete(do_nothing)" json:"-"`
	Model
}

func NewCourse() *Course {
	return &Course{}
}

//获取单周课表
func (this *Course) GetCourse() (data []Course, err error) {
	//登录CAS
	c, err := this.Classes.Student.LoginCAS()
	if err != nil {
		return nil, err
	}
	//获取课表
	result, err := course.Get(this.Term, this.Weekly, c)
	if err != nil {
		return nil, err
	}
	var courseData []Course
	for _, v := range result {
		course := Course{
			Term:     this.Term,
			Weekly:   this.Weekly,
			Name:     v.Name,
			Position: v.Position,
			Class:    v.Class,
			Teacher:  v.Teacher,
			Classes:  this.Classes,
		}
		//判断是否课程是否保存本地
		if err = DB().QueryTable(GetTable("course")).Filter("Term", this.Term).Filter("Weekly", this.Weekly).Filter("Name", v.Name).Filter("Position", v.Position).Filter("Class", v.Class).Filter("Teacher", v.Teacher).One(this); err != nil {
			courseData = append(courseData, course)
		}
	}
	if _, err = DB().InsertMulti(100, courseData); err != nil {
		return nil, errors.New("获取课程表失败123")
	}
	if num, _ := DB().QueryTable(GetTable("course")).Filter("Classes__Id", this.Classes.Id).Filter("Term", this.Term).Filter("Weekly", this.Weekly).All(&data, "Id", "Weekly", "Name", "Position", "Class", "Teacher"); num == 0 {
		return nil, errors.New("课程表查询失败")
	}
	return data, nil
}

func (this *Course) GetCourseHost() (data []Course, err error) {
	//获取学生学号密码
	student := NewStudent()
	if err = DB().QueryTable(GetTable("student")).Filter("User__Id", this.Classes.Student.User.Id).One(student); err != nil {
		return nil, errors.New("未绑定学号")
	}
	this.Classes = student.Classes
	this.Classes.Student = student
	if num, _ := DB().QueryTable(GetTable("course")).Filter("Classes__Id", this.Classes.Id).Filter("Term", this.Term).Filter("Weekly", this.Weekly).All(&data, "Id", "Weekly", "Name", "Position", "Class", "Teacher"); num == 0 {
		return nil, errors.New("课程表查询失败")
	}
	return data, nil
}

//获取下节课
func (this *Course) GetNextClass() (err error) {
	if err = DB().QueryTable(GetTable("course")).Filter("Classes__Id", this.Classes.Id).Filter("Term", this.Term).Filter("Weekly", this.Weekly).Filter("Position", this.Position).One(this); err != nil {
		return err
	}
	return nil
}
