package main

import (
	"github.com/micro/go-micro/v2"
	"micro-project/proto/course"
	"micro-project/services"
)

func main() {
	service := micro.NewService(micro.Name("go.micro.api.zyxm.course"))

	service.Init()

	err := course.RegisterCourseServiceHandler(service.Server(), new(services.CourseServiceImpl))
	if err != nil {
		panic(err)
	}

	err = service.Run()
	if err != nil {
		panic(err)
	}
}
