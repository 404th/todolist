package repository

import (
	"github.com/404th/todolist/model"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	User
	Todo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(user model.SignInInput) (model.User, error)
}

type User interface{}

type Todo interface{}
