package models

import (
	"errors"
	"fjsdxy-plus/helper"
	"github.com/pig0224/fjsdxy/jwc"
	"github.com/pig0224/fjsdxy/jwc/week"
	"strconv"
	"time"
)

type Week struct {
	Id     int          `orm:"auto;pk;index" json:"id"`
	Term   string       `orm:"size(30)" json:"term"`
	Weekly uint64       `orm:"unique" json:"weekly"`
	Week   time.Weekday `orm:"unique" json:"week"`
	Today  string       `orm:"type(date)" json:"today"`
	Model
}

func NewWeek() *Week {
	return &Week{}
}

func (this *Week) Pull() (err error) {
	//登录教务处
	c, err := jwc.Login("1380180058", "244538")
	if err != nil {
		return errors.New("登录教务处失败")
	}
	//获取学校周历数据
	result, err := week.Get(this.Term, c)
	if err != nil {
		return errors.New("获取周历数据失败")
	}
	var weekData []Week
	var data Week
	for _, v := range result {
		data = Week{
			Term:   this.Term,
			Weekly: v.Weekly,
			Week:   v.Week,
			Today:  v.Today,
		}
		weekData = append(weekData, data)
	}
	nums, err := DB().InsertMulti(100, weekData)
	if err != nil {
		return errors.New("保存周历数据失败")
	}
	helper.Logger.Info("周历保存：" + strconv.FormatInt(nums, 10) + "条")
	return nil
}

func (this *Week) GetWeek() (err error) {
	if err = DB().Read(this, "Today"); err != nil {
		return errors.New("查询失败")
	}
	return nil
}
