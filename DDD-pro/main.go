package main

import (
	"eventbus/domain/models/course"
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	c := course.New(
		course.WithCourseName("name"),
		course.WithCourseImage("image"),
		course.WithCourseTime(3600),
	)
	c.CourseTime.AddSecond(100)
	b, _ := jsoniter.Marshal(c)
	fmt.Println(string(b))
}
