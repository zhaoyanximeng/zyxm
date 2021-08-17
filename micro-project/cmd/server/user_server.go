package main

import (
	"github.com/micro/go-micro/v2"
	g "micro-project/handler/grpc"
	"micro-project/proto/course"
	"micro-project/proto/user"
)

func main() {
	service := micro.NewService(micro.Name("go.micro.api.zyxm.user"))

	service.Init()

	course.NewCourseService("go.micro.api.zyxm.user", service.Client())

	err := user.RegisterUserServiceHandler(service.Server(), g.NewUserService(service.Client()))
	if err != nil {
		panic(err)
	}

	err = service.Run()
	if err != nil {
		panic(err)
	}
}
