package repository

import (
	"github.com/404th/todolist/model"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}

var (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

func (ar AuthRepository) CreateUser(user model.User) (int, error) {
	var id int = 0

	query_string := `
		INSERT INTO users (
			username,
			name,
			password_hash
		) VALUES (
			$1,
			$2,
			$3
		) RETURNING id
	`
	row := ar.db.QueryRow(query_string, user.Username, user.Name, user.Password)
	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (ar AuthRepository) GetUser(user model.SignInInput) (model.User, error) {
	var resp_user model.User

	qs := `SELECT id FROM users WHERE username=$1 AND password_hash=$2`
	if err := ar.db.Get(&resp_user, qs, user.Username, user.Password); err != nil {
		return resp_user, err
	}

	return resp_user, nil
}
