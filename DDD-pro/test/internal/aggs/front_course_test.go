package aggs

import (
	"eventbus/domain/aggs/front_course"
	"eventbus/domain/models/course_model"
	"eventbus/domain/models/course_price"
	"eventbus/infrastructure/mysql_data"
	"eventbus/test/internal/conf"
	"testing"
)

// 测试商品聚合查询
func TestFrontCourseAgg_QueryDetail(t *testing.T) {
	cmr := mysql_data.NewCourseMainRepo(conf.DB)
	cpr := mysql_data.NewCoursePriceRepo(conf.DB)

	cm := course_model.New()
	cm.CourseID = 3
	cp := course_price.New(course_price.WithCourseID(cm.CourseID))

	fc := front_course.NewFrontCourseAgg(cm,cp,cmr,cpr)
	err := fc.QueryDetail()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n",fc.CourseMain)
	t.Logf("%+v\n",fc.CoursePrice)
	t.Logf("%+v\n",fc.CoursePrice.Price)
}