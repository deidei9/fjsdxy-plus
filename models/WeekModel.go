package models

import "time"

type Week struct {
	Id     int       `orm:"auto;pk;index" json:"id"`
	Term   string    `orm:"size(30)" json:"term"`
	Weekly int       `orm:"unique" json:"weekly"`
	Week   int       `orm:"unique" json:"week"`
	Today  time.Time `orm:"type(date)" json:"today"`
	Model
}

func NewWeek() *Week {
	return &Week{}
}
