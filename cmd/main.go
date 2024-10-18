package main

import (
	"log"

	"github.com/DavidG9999/my_test_app/interal/handler"
	"github.com/DavidG9999/my_test_app/interal/repository"
	srv "github.com/DavidG9999/my_test_app/interal/server"
	"github.com/DavidG9999/my_test_app/interal/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(srv.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
