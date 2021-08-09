package repos

type CourseMainRepo interface {
	FindByID(m Model) error
}