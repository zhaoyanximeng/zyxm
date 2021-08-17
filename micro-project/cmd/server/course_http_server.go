package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
	"log"
	h "micro-project/handler/http"
)

func main() {

	r := gin.New()
	r.GET("/test", h.CourseTest)

	service := web.NewService(
		web.Name("go.micro.api.zyxm.http.course"),
		web.Handler(r),
	)

	service.Init()
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
