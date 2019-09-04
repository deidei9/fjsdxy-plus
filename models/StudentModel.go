package models

import (
	"errors"
	"github.com/gocolly/colly"
	sy "github.com/pig0224/fjsdxy"
	"github.com/pig0224/fjsdxy/cas"
	"github.com/pig0224/fjsdxy/jwc"
	"github.com/pig0224/fjsdxy/jwc/student"
	"github.com/pig0224/fjsdxy/xg"
	//"github.com/pig0224/fjsdxy/jwc/student"
)

type Student struct {
	Id          int      `orm:"auto;pk;index" json:"id"`
	StudentId   string   `orm:"size(30)" json:"student_id"`
	Password    string   `orm:"size(30)" json:"password"`
	StudentName string   `orm:"size(30)" json:"student_name"`
	IsNotice    float64  `orm:"digits(1);decimals(0);default(0)" json:"isNotice"`
	User        *User    `orm:"rel(one)" json:"user_info"`
	Classes     *Classes `orm:"rel(one)" json:"classes_info"`
	Exam        []*Exam  `orm:"reverse(many)" json:"exam_info"`
	Model
}

func NewStudent() *Student {
	return &Student{}
}

//登录SSO
func (this *Student) LoginSSO() (c *colly.Collector, err error) {
	if c, err = sy.SSO_Login(this.StudentId, this.Password); err != nil {
		return nil, errors.New("校务系统出错，请稍后再试")
	}
	return c, nil
}

//登录CAS
func (this *Student) LoginCAS() (c *colly.Collector, err error) {
	if c, err = cas.Login(this.StudentId, this.Password); err != nil {
		return nil, errors.New("校务系统出错，请稍后再试")
	}
	return c, nil
}

//登录教务处
func (this *Student) LoginJWC() (c *colly.Collector, err error) {
	if c, err = jwc.Login(this.StudentId, this.Password); err != nil {
		return nil, errors.New("校务系统出错，请稍后再试")
	}
	return c, nil
}

//登录学工处
func (this *Student) LoginXG() (c *colly.Collector, err error) {
	if c, err = xg.Login(this.StudentId, this.Password); err != nil {
		return nil, errors.New("校务系统出错，请稍后再试")
	}
	return c, nil
}

//绑定学号
func (this *Student) Bind() (err error) {
	if err := DB().Read(this.User); err != nil {
		return errors.New("未登录")
	}
	//判断用户来源
	userFrom := this.User.UserFrom()
	oldStudent := NewStudent()
	if err = DB().QueryTable(GetTable("student")).Filter("StudentId", this.StudentId).RelatedSel().One(oldStudent); err == nil {
		//判断已绑定的用户来源
		oldUserFrom := oldStudent.User.UserFrom()
		if userFrom == oldUserFrom { //重复绑定
			return errors.New("学号已被绑定")
		}
	}
	if _, err = this.LoginSSO(); err != nil {
		return err
	}
	//获取用户信息
	if err = this.BindInfo(); err != nil {
		return err
	}
	//学号和用户绑定
	if err = DB().QueryTable(GetTable("student")).Filter("User__Id", this.User.Id).RelatedSel().One(this); err != nil {
		//第一次绑定
		if _, err = DB().Insert(this); err != nil {
			return errors.New("绑定失败")
		}
	} else {
		//更新绑定
		if _, err = DB().Update(this); err != nil {
			return errors.New("更新绑定失败")
		}
	}
	return nil
}

//解绑学号
func (this *Student) UnBind() (err error) {
	if err = DB().QueryTable(GetTable("student")).Filter("User__Id", this.User.Id).One(this); err != nil {
		return errors.New("未绑定学号")
	}
	if _, err := DB().Delete(this); err != nil {
		return errors.New("解绑失败")
	}
	return nil
}

//绑定学生信息&更新学生信息
func (this *Student) BindInfo() (err error) {
	//登录教务处
	c, err := this.LoginJWC()
	if err != nil {
		return err
	}
	//获取学生信息
	info, err := student.Get(c)
	if err != nil {
		return err
	}
	//班级信息不存在，创建
	this.Classes = NewClasses()
	this.Classes.ClassName = info.ClassName
	if err = DB().Read(this.Classes, "ClassName"); err != nil {
		//班级不存在，创建班级
		this.Classes.ClassName = info.ClassName
		this.Classes.Major = info.Major
		this.Classes.College = info.College
		if _, err := DB().Insert(this.Classes); err != nil {
			return errors.New("班级信息创建错误")
		}
	} else {
		//更新学生信息
		this.Classes.ClassName = info.ClassName
		this.Classes.Major = info.Major
		this.Classes.College = info.College
		if _, err := DB().Update(this.Classes); err != nil {
			return errors.New("班级信息更新错误")
		}
	}
	this.StudentName = info.StudentName
	return nil
}

//获取学生信息
func (this *Student) GetInfo() (err error) {
	if err := DB().Read(this.User); err != nil {
		return errors.New("未登录")
	}
	if err = DB().QueryTable(GetTable("student")).Filter("User__Id", this.User.Id).RelatedSel().One(this); err != nil {
		return errors.New("未绑定学号")
	}
	return nil
}

//更新学生信息
func (this *Student) PullInfo() (err error) {
	_ = DB().QueryTable(GetTable("student")).Filter("User__Id", this.User.Id).One(this)
	_ = this.BindInfo() //更新数据
	if _, err = DB().Update(this); err != nil {
		return err
	}
	return nil
}

//设置上课通知 0：开启  1：关闭
func (this *Student) ClassNotice() (err error) {
	if err := this.GetInfo(); err != nil {
		return err
	}
	if this.IsNotice == 0 {
		this.IsNotice = 1
	} else {
		this.IsNotice = 0
	}
	if _, err = DB().Update(this); err != nil {
		return errors.New("设置失败")
	}
	return nil
}
