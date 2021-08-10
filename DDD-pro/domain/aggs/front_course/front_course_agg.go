package front_course

import (
	"eventbus/domain/models/course_model"
	"eventbus/domain/models/course_price"
	"eventbus/domain/models/repos"
	"eventbus/infrastructure/errorstring"
)

// 前台商品（聚合）的重要元素
type FrontCourseAgg struct {
	CourseMain *course_model.CourseMain	// 聚合根
	CoursePrice *course_price.CoursePrice // 商品价格实体

	CourseMainRepo repos.CourseMainRepo // 仓储
	CoursePriceReop repos.CoursePriceRepo // 仓储
}

func NewFrontCourseAgg(cm *course_model.CourseMain, cp *course_price.CoursePrice,
	cmr repos.CourseMainRepo, cpr repos.CoursePriceRepo) *FrontCourseAgg {
	if cm == nil {
		panic("root course main err")
	}

	fc := &FrontCourseAgg{CourseMain: cm, CoursePrice: cp, CourseMainRepo: cmr, CoursePriceReop: cpr}
	fc.CourseMain.Repo = cmr
	fc.CoursePrice.Repo = cpr

	return fc
}

func (fca *FrontCourseAgg) QueryDetail() error {
	if fca.CourseMain.CourseID <= 0 {
		//return fmt.Errorf("course id err")
		return errorstring.NewNotFoundIDError("course_main")
	}

	// 取出商品详细信息
	err := fca.CourseMain.Load()
	if err != nil {
		//return errors.New(fmt.Sprintf("get course main data err:%s",err))
		return errorstring.NewNotFoundDataError("course_main",err.Error())
	}

	// 取出商品价格详细信息
	err = fca.CoursePrice.Load()
	if err != nil {
		//return errors.New(fmt.Sprintf("get course price data err:%s",err))
		return errorstring.NewNotFoundDataError("course_price",err.Error())
	}

	return nil
}

func (fca *FrontCourseAgg) CreateFrontCourse() error {
	if err := fca.CourseMain.New(); err != nil {
		return errorstring.NewModelCreateError("course_main",err.Error())
	}
	fca.CoursePrice.CourseID = fca.CourseMain.CourseID
	if err := fca.CoursePrice.New(); err != nil {
		return errorstring.NewModelCreateError("course_price",err.Error())
	}

	return nil
}