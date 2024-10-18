package main

import (
	"log"

	"github.com/DavidG9999/my_test_app/interal/handler"
	srv "github.com/DavidG9999/my_test_app/interal/server"
)

func main() {
	handler:=new(handler.Handler)
	srv := new(srv.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
