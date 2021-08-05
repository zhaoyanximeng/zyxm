package course

type CourseInfo struct {
	CourseName  string `json:"course_name" gorm:"column:course_name"`
	CourseImage string `json:"course_image" gorm:"column:course_image"`
}

func NewCourseInfo() *CourseInfo {
	return &CourseInfo{}
}

func WithCourseName(name string) CourseMainOption {
	return func(c *CourseMain) {
		c.CourseInfo.CourseName = name
	}
}

func WithCourseImage(image string) CourseMainOption {
	return func(c *CourseMain) {
		c.CourseInfo.CourseImage = image
	}
}
