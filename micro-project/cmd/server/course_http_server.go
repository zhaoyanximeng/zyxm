package main

import (
	"context"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/web"
	"log"
	"micro-project/proto/course"
	"net/http"
)

func main() {
	service := web.NewService(web.Name("go.micro.api.zyxm.http.course"))

	client := grpc.NewClient()
	service.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		cs := course.NewCourseService("go.micro.api.zyxm.course", client)
		rsp, _ := cs.ListForTop(context.Background(), &course.ListRequest{Size: 10})
		log.Println(rsp.Result)
		writer.Write([]byte("http api test"))
	})

	service.Init()
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
