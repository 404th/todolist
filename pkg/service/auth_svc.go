package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/404th/todolist/model"
	"github.com/404th/todolist/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

var salt string = "my super secret sall for hashing password!"
var signingKey string = "ytgsdubhfinkfmsdflsdmfosdflnsodjfksnf"
var tokenTTL = 12 * time.Hour

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

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

func (as *AuthService) GetUser(user model.SignInInput) (string, error) {
	user.Password = hashPassword(user.Password)
	resp_user, err := as.repo.GetUser(user)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			resp_user.ID,
		},
	)

	return token.SignedString([]byte(signingKey))
}

func hashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
