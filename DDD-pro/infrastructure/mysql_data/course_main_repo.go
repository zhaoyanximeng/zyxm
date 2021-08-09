package mysql_data

import (
	"eventbus/domain/models/course_model"
	"eventbus/domain/models/repos"
	"gorm.io/gorm"
)

type CourseMainRepo struct {
	db *gorm.DB
}

func NewCourseMainRepo(db *gorm.DB) *CourseMainRepo {
	return &CourseMainRepo{db: db}
}

func (cmr *CourseMainRepo)FindByID(m repos.Model) error {
	return cmr.db.Table("courses").Where("course_id=?",
		m.(*course_model.CourseMain).CourseID).First(m).Error
}