package models

import (
	"eventbus/domain/models/course"
	"eventbus/test/internal/conf"
	"testing"
)

func CourseCreate(name string,time int64) error {
	c := course.New(
		course.WithCourseName(name),
		course.WithCourseTime(time),
	)

	return conf.DB.Table("courses").Create(c).Error
}

func TestCourseCreate(t *testing.T) {
	data := [] struct{
		name string
		time int64
	} {
		{
			name:"golang",
			time:233,
		},{
			name:"c",
			time:1232,
		},{
			name:"c++",
			time:2343,
		},
	}
	for i := range data {
		name,time := data[i].name,data[i].time
		err := CourseCreate(name,time)
		if err != nil {
			t.Errorf("课程%s创建出错，err:%s\n",data[i].name,err)
		}
	}
}