package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"micro-project/proto/user"
)

type UserService struct {

}

func (u *UserService) Test(ctx context.Context, in *user.UserRequest, out *user.UserResponse) error {
	out.Ret = "user" + in.Id
	return nil
}

func main() {
	service := micro.NewService(micro.Name("go.micro.api.zyxm.user"))

	service.Init()

	err := user.RegisterUserServiceHandler(service.Server(),new(UserService))
	if err != nil {
		panic(err)
	}

	err = service.Run()
	if err != nil {
		panic(err)
	}
}