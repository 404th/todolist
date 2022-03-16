package service

import (
	"github.com/404th/todolist/model"
	"github.com/404th/todolist/pkg/repository"
)

type Service struct {
	Authorization
	User
	Todo
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}

type Authorization interface {
	CreateUser(model.User) (int, error)
}

type User interface{}

type Todo interface{}
