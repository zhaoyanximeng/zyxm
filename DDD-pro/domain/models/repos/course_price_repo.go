package repos

type CoursePriceRepo interface {
	FindByID(m Model) error
}
