package model

type User struct {
	ID       int    `json:"-"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
