package front_course

import (
	"eventbus/domain/models/course_model"
	"eventbus/domain/models/course_price"
)

// 前台商品（聚合）的重要元素
type FrontCourseAgg struct {
	CourseMain *course_model.CourseMain	// 聚合根
	CoursePrice *course_price.CoursePrice // 商品价格实体
}

func (fca *FrontCourseAgg) QueryDetail() {

}