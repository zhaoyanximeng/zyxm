package main

import (
	"eventbus/domain/models/course_model"
	"eventbus/domain/models/course_price"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"time"
)

func main() {
	c := course_model.New(
		course_model.WithCourseName("name"),
		course_model.WithCourseImage("image"),
		course_model.WithCourseTime(3600),
	)
	c.CourseTime.AddSecond(100)
	b, _ := jsoniter.Marshal(c)
	fmt.Println(string(b))

	cp := course_price.New(
		course_price.WithMarketPrice(20),
		course_price.WithSalePrice(40),
		course_price.WithComment("comment"),
		course_price.WithSetdata(time.Now()))
	bc,_ := jsoniter.Marshal(cp)
	fmt.Println(string(bc))
}
