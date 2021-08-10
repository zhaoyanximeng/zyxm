package aggs

import (
	"eventbus/domain/aggs/front_course"
	"eventbus/domain/models/course_model"
	"eventbus/domain/models/course_price"
	"eventbus/infrastructure/mysql_data"
	"eventbus/test/internal/conf"
	"fmt"
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
		t.Fatal(err)
	}
	t.Logf("%+v\n",fc.CourseMain)
	t.Logf("%+v\n",fc.CoursePrice)
	t.Logf("%+v\n",fc.CoursePrice.Price)
}

func TestFrontCourseAgg_CreateFrontCourse(t *testing.T) {
	tx := conf.DB.Begin()
	cmr := mysql_data.NewCourseMainRepo(tx)
	cpr := mysql_data.NewCoursePriceRepo(tx)

	cm := course_model.New(course_model.WithCourseName("云原生"))
	cm.CourseID = 12

	cp := course_price.New(
		course_price.WithMarketPrice(50),
		course_price.WithSalePrice(45))

	fc := front_course.NewFrontCourseAgg(cm,cp,cmr,cpr)
	err := fc.CreateFrontCourse()
	if err != nil {
		tx.Rollback()
		t.Fatal(err)
	} else {
		tx.Commit()
	}

	fmt.Printf("%+v\n",fc.CourseMain)
	fmt.Printf("%+v\n",fc.CoursePrice)
	fmt.Printf("%+v\n",fc.CoursePrice.Price)
}