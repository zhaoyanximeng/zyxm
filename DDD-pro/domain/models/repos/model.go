package repos

type Model interface {
	Name() string
	Load() error
}
