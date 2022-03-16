package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/404th/todolist/model"
	"github.com/404th/todolist/pkg/repository"
)

var salt string = "my super secret sall for hashing password!"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (as *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = hashPassword(user.Password)
	return as.repo.CreateUser(user)
}

func hashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
