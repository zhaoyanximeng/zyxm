package main

import (
	"github.com/micro/go-micro/v2"
	"micro-project/proto/course"
)

func main() {
	service := micro.NewService(micro.Name("go.micro.api.zyxm.course"))

	service.Init()

	err := course.RegisterCourseServiceHandler(service.Server(), new(course.CourseServiceImpl))
	if err != nil {
		panic(err)
	}

	err = service.Run()
	if err != nil {
		panic(err)
	}
}
