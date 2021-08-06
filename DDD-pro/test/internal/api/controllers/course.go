package controllers

import (
	"eventbus/domain/models/course"
	"eventbus/test/internal/conf"
	"eventbus/test/internal/utils/errcode"
	"eventbus/test/internal/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CourseControllers struct {

}

func (cc *CourseControllers) Name() string {
	return "course_controllers"
}

func (cc *CourseControllers) Add(c *gin.Context) {
	co := course.New()
	if err := c.ShouldBind(co);err != nil {
		logrus.Errorf("bind err:%s",err)
		response.Error(c,errcode.ErrInvalidParam.WithDetails(err.Error()))
	}
	err := conf.DB.Table("courses").Create(co).Error
	if err != nil {
		logrus.Errorf("query data err:%s",err)
		response.Error(c,errcode.ErrInternalServer.WithDetails(err.Error()))
	}

	response.Success(c,gin.H{
		"course_id":co.CourseID,
	})
}

func (cc *CourseControllers) BuildRoute(r *gin.RouterGroup) {
	r.POST("/courses",cc.Add)
}