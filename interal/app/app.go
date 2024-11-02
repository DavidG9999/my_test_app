package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	httpsrv "github.com/DavidG9999/my_test_app/interal/app/httpserver"
	"github.com/DavidG9999/my_test_app/interal/handler"
	"github.com/DavidG9999/my_test_app/interal/repository"
	"github.com/DavidG9999/my_test_app/interal/repository/postgres"
	"github.com/DavidG9999/my_test_app/interal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type App struct {
	Repos      *repository.Repository
	Services   *service.Service
	Handlers   *handler.Handler
	HTTPServer *http.Server
}

func RunApp() {

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to connect db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	httpsrv := new(httpsrv.Server)
	go func() {
		if err := httpsrv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App shutting down")

	if err := httpsrv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

}
