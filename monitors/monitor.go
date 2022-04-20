package monitors

type Monitor[T any] interface {
	Monitor() (bool, T)
}
