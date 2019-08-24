package models

import (
	"errors"
	"github.com/pig0224/fjsdxy/jwc"
	"github.com/pig0224/fjsdxy/jwc/exam"
)

type Exam struct {
	Id      int      `orm:"auto;pk;index" json:"id"`
	Term    string   `orm:"size(30)" json:"term"`
	Name    string   `orm:"size(30)" json:"name"`
	Type    string   `orm:"size(30)" json:"type"`
	Score   string   `orm:"size(5)" json:"score"`
	Credits string   `json:"credits"`
	Student *Student `orm:"rel(fk)" json:"-"`
	Model
}

func NewExam() *Exam {
	return &Exam{}
}

//根据学生获取成绩
func (this *Exam) GetByStu() (num int64, err error) {
	num, err = DB().QueryTable(GetTable("exam")).Filter("Student__Id", this.Student.Id).All(this)
	if err != nil {
		return num, errors.New("获取失败")
	}
	return num, nil
}

//获取学生成绩数据
func (this *Exam) PullExam() (num int64, err error) {
	studentId := this.Student.StudentId
	pass := this.Student.Password
	c, err := jwc.Login(studentId, pass)
	if err != nil {
		return 0, err
	}
	result, err := exam.Get("", c)
	if err != nil {
		return 0, err
	}
	//写入数据
	var examData []Exam
	var data Exam
	for _, v := range result {
		//判断成绩是否存在
		if err := DB().QueryTable(GetTable("exam")).Filter("Student__id", this.Student.Id).Filter("Name", v.Name).Filter("Score", v.Score).RelatedSel().One(this); err != nil {
			data = Exam{
				Student: this.Student,
				Term:    v.Term,
				Name:    v.Name,
				Type:    v.Type,
				Score:   v.Score,
				Credits: v.Credits,
			}
			examData = append(examData, data)
		}
	}
	num, _ = DB().InsertMulti(100, examData)
	return num, nil
}
