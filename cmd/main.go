package main

import (
	"log"

	"github.com/404th/todolist"
	"github.com/404th/todolist/pkg/handler"
	"github.com/404th/todolist/pkg/repository"
	"github.com/404th/todolist/pkg/service"
	_ "github.com/lib/pq"
)

func main() {
	// init Config
	// if err := initConfig(); err != nil {
	// 	log.Fatalf("Error while initializing Config file: %s", err.Error())
	// }

	// repos := repository.NewRepository()

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5555",
		Username: "postgres",
		DBName:   "simple_bank",
		Password: "pass",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	// initialize Server
	srv := new(todolist.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while runnig http server: %s", err.Error())
	}
}
