package mysql_data

import (
	"eventbus/domain/models/course_price"
	"eventbus/domain/models/repos"
	"fmt"
	"gorm.io/gorm"
)

type CoursePriceRepo struct {
	db *gorm.DB
}

func NewCoursePriceRepo(db *gorm.DB) *CoursePriceRepo {
	return &CoursePriceRepo{db: db}
}

func (cpr *CoursePriceRepo) FindByID(m repos.Model) error {
	_,ok := m.(*course_price.CoursePrice)
	if !ok {
		fmt.Println("δΈει")
	}

	return cpr.db.Table("course_prices").Where("course_id = ? and iscurrent = 1",
		m.(*course_price.CoursePrice).CourseID).First(m).Error
}

func (cpr *CoursePriceRepo) New(m repos.Model) error {
	return cpr.db.Table("course_prices").Create(m).Error
}
