package main

import (
	"log"

	srv "github.com/DavidG9999/my_test_app/interal/server"
)

func main() {
	srv := new(srv.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
