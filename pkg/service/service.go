package service

type Authentication interface {
}

type User interface {
}

type Todo interface {
}

type Service struct {
	Authentication
	User
	Todo
}

func NewService() *Service {
	return &Service{}
}
