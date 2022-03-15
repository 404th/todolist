package repository

type Authentication interface {
}

type User interface {
}

type Todo interface {
}

type Repository struct {
	Authentication
	User
	Todo
}

func NewRepository() *Repository {
	return &Repository{}
}
