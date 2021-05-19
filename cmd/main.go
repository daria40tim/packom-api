package main

import (
	"github.com/daria40tim/packom"
	"github.com/daria40tim/packom/pkg/handler"
	"github.com/daria40tim/packom/pkg/repository"
	"github.com/daria40tim/packom/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "5770",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(packom.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatal("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
