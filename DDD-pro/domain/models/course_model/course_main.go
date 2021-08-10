package course_model

import "eventbus/domain/models/repos"

type CourseMain struct {
	CourseID      int                  `json:"course_id" gorm:"column:course_id;AUTO_INCREMENT;PRIMARY_KEY"`
	CourseInfo    *CourseInfo          `json:"course_info" gorm:"embedded"` // 课程信息
	CourseTime    *CourseTime          `json:"course_time" gorm:"embedded"` // 课程时长
	CourseAddDate *CourseAddDate       `json:"course_adddate" gorm:"embedded"` // 入库时间
	Repo          repos.CourseMainRepo `gorm:"-"`
}

func (*CourseMain) Name() string {
	return "course_main"
}

func (cm *CourseMain) Load() error {
	return cm.Repo.FindByID(cm)
}

func (cm *CourseMain) New() error {
	return cm.Repo.New(cm)
}

func WithRepo(repo repos.CourseMainRepo) CourseMainOption {
	return func(c *CourseMain) {
		c.Repo = repo
	}
}

func New(attrs ...CourseMainOption) *CourseMain {
	c := &CourseMain{
		CourseInfo:    NewCourseInfo(),
		CourseTime:    NewCourseTime(),
		CourseAddDate: NewCourseAddDate(),
	}
	CourseMainOptions(attrs).apply(c)

	return c
}

type CourseMainOption func(model *CourseMain)

type CourseMainOptions []CourseMainOption

func (c CourseMainOptions) apply(u *CourseMain) {
	for _, f := range c {
		f(u)
	}
}
