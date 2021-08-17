package grpc

import (
	"context"
	"github.com/micro/go-micro/v2/client"
	"log"
	"micro-project/proto/course"
	"micro-project/proto/user"
)

type UserService struct {
	client client.Client
}

func NewUserService(client client.Client) *UserService {
	return &UserService{client: client}
}

func (u *UserService) Test(ctx context.Context, in *user.UserRequest, out *user.UserResponse) error {
	out.Ret = "user" + in.Id
	cs := course.NewCourseService("go.micro.api.zyxm.course", u.client)
	rsp, _ := cs.ListForTop(ctx, &course.ListRequest{Size: 10})
	log.Println(rsp.Result)
	return nil
}
