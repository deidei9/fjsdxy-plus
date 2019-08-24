package helper

import (
	"fmt"
	"strconv"
	"time"
)

func init() {
	//初始化Logs
	InitLogs()
	//初始化Cache
	InitCache()
}

//Time转换课程表position
func Time2Pos(week, hour, min int) (position []string) {
	thisWeek := strconv.Itoa(week + 1)

	if hour < 8 {
		position = []string{thisWeek + "AB", thisWeek + "CD", thisWeek + "EF", thisWeek + "GH", thisWeek + "IJ"}
		return position
	}

	if hour == 8 && min < 0 {
		position = []string{thisWeek + "AB", thisWeek + "CD", thisWeek + "EF", thisWeek + "GH", thisWeek + "IJ"}
		return position
	}

	if hour < 10 {
		position = []string{thisWeek + "CD", thisWeek + "EF", thisWeek + "GH", thisWeek + "IJ"}
		return position
	}
	if hour == 10 && min < 5 {
		position = []string{thisWeek + "CD", thisWeek + "EF", thisWeek + "GH", thisWeek + "IJ"}
		return position
	}

	if hour < 14 {
		position = []string{thisWeek + "EF", thisWeek + "GH", thisWeek + "IJ"}
		return position
	}
	if hour == 14 && min < 30 {
		position = []string{thisWeek + "EF", thisWeek + "GH", thisWeek + "IJ"}
		return position
	}

	if hour < 16 {
		position = []string{thisWeek + "GH", thisWeek + "IJ"}
		return position
	}
	if hour == 16 && min < 15 {
		position = []string{thisWeek + "GH", thisWeek + "IJ"}
		return position
	}

	if hour < 19 {
		position = []string{thisWeek + "IJ"}
		return position
	}
	if hour == 19 && min < 0 {
		position = []string{thisWeek + "IJ"}
		return position
	}
	return nil
}

//获取上课时间
func GetClassTime(postion string) (beginTime string) {
	p := postion[1:len(postion)]
	fmt.Println(p)
	switch p {
	case "AB":
		beginTime = "8:00-9:35"
	case "CD":
		beginTime = "10:05-11:40"
	case "EF":
		beginTime = "14:30-16:05"
	case "GH":
		beginTime = "16:15-17:50"
	case "IJ":
		beginTime = "19:00-20:35"
	default:
		beginTime = ""
	}
	return beginTime
}

//获取星期几
func GetWeek(weekday time.Weekday) (week string) {
	switch weekday {
	case 0:
		week = "星期一"
	case 1:
		week = "星期二"
	case 2:
		week = "星期三"
	case 3:
		week = "星期四"
	case 4:
		week = "星期五"
	case 5:
		week = "星期六"
	case 6:
		week = "星期天"
	default:
		week = ""
	}
	return week
}
