package main

import (
	"log"

	quotes_memorizer "github.com/Twofold-One/quotes-memorizer-api"
	"github.com/Twofold-One/quotes-memorizer-api/pkg/handler"
	"github.com/Twofold-One/quotes-memorizer-api/pkg/repository"
	"github.com/Twofold-One/quotes-memorizer-api/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(quotes_memorizer.Server)

	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigFile("config")
	return viper.ReadInConfig()
}