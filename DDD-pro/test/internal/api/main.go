package main

import (
	"eventbus/test/internal/api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	v1 := r.Group("/test/v1")
	{
		new(controllers.CourseControllers).BuildRoute(v1)
	}
	r.Run(":81")
}
