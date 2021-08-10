package repos

type CoursePriceRepo interface {
	FindByID(m Model) error
	New(m Model) error
}
