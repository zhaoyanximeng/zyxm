package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client/grpc"
	"micro-project/proto/course"
	"net/http"
)

func CourseTest(c *gin.Context) {
	client := grpc.NewClient()
	cs := course.NewCourseService("go.micro.api.zyxm.course", client)
	rsp, _ := cs.ListForTop(context.Background(), &course.ListRequest{Size: 10})
	c.JSON(http.StatusOK, gin.H{"result": rsp.Result})
}
