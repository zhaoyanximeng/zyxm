package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"micro-project/proto/user"
)

type UserService struct {

}

func (u *UserService) Test(ctx context.Context, in *user.UserRequest, out *user.UserResponse) error {
	out.Ret = "user" + in.Id
	return nil
}

func main() {
	etcdRegistry := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:22379"}
	})

	service := micro.NewService(micro.Name("go.micro.api.zyxm.user"),
		micro.Registry(etcdRegistry),
		)

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