package repository

import (
	"fmt"

	"github.com/404th/todolist/config"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func NewPsqlDB(cfg config.DBConfig) *sqlx.DB {
	psql_s := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode,
	)
	db, err := sqlx.Open("postgres", psql_s)
	if err != nil {
		logrus.Fatalf("Failed to create database: %s", err.Error())
		return nil
	}

	if err = db.Ping(); err != nil {
		logrus.Fatalf("Failed while checking for ping: %s", err.Error())
		return nil
	}

	return db
}
