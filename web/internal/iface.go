package internal

type Server interface {
	Listen() error
	Shutdown() error
}
