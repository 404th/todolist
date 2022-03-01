package main

import (
	"log"

	"github.com/404th/todolist"
)

func main() {
	// initialize Server
	srv := new(todolist.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error while runnig http server: %s", err.Error())
	}

}
