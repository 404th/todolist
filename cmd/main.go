package main

import (
	"github.com/404th/todolist/config"
	"github.com/404th/todolist/pkg/handler"
	"github.com/404th/todolist/pkg/repository"
	"github.com/404th/todolist/pkg/service"
	"github.com/404th/todolist/server"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	new_server := new(server.Server)
	cfg := config.DBConfig{
		Host:     "localhost",
		Port:     "6565",
		User:     "postgres",
		DBName:   "simple_bank",
		Password: "postgres",
		SSLMode:  "disable",
	}

	new_db := repository.NewPsqlDB(cfg)
	new_repo := repository.NewRepository(new_db)
	new_service := service.NewService(new_repo)
	new_handler := handler.NewHandler(new_service)

	server_port := "7777"
	if err := new_server.RunServer(server_port, new_handler.InitRoutes()); err != nil {
		logrus.Fatalf("Failed to run server: %s", err.Error())
		return
	}
}
