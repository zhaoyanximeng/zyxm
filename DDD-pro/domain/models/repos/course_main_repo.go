package repos

type CourseMainRepo interface {
	FindByID(m Model) error
	New(m Model) error
}