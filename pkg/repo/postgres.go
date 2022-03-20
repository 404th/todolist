package repo

import (
	"fmt"

	"github.com/404th/todolist/configs"
	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(cfg configs.DB_config) (*sqlx.DB, error) {
	new_db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User_Name, cfg.DB_Name, cfg.Password, cfg.SSLMode,
		),
	)
	if err != nil {
		return nil, err
	}

	if err := new_db.Ping(); err != nil {
		return nil, err
	}

	return new_db, nil
}
