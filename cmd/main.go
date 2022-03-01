package main

import (
	"log"

	"github.com/404th/todolist"
	"github.com/404th/todolist/pkg/handler"
	"github.com/404th/todolist/pkg/repository"
	"github.com/404th/todolist/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	// init Config
	if err := initConfig(); err != nil {
		log.Fatalf("Error while initializing Config file: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	// initialize Server
	srv := new(todolist.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while runnig http server: %s", err.Error())
	}
}

// init config/viper file
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
