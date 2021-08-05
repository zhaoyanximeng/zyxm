package controllers

import (
	"eventbus/domain/models/course"
	"eventbus/test/internal/conf"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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
		c.JSON(http.StatusBadRequest,gin.H{
			"err":err,
		})
	}
	err := conf.DB.Table("courses").Create(co).Error
	if err != nil {
		logrus.Errorf("bind err:%s",err)
		c.JSON(http.StatusInternalServerError,gin.H{
			"err":err,
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"course_id":co.CourseID,
	})
}

func (cc *CourseControllers) BuildRoute(r *gin.RouterGroup) {
	r.POST("/courses",cc.Add)
}