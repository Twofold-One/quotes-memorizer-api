package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	quotes_memorizer "github.com/Twofold-One/quotes-memorizer-api"
	"github.com/Twofold-One/quotes-memorizer-api/pkg/handler"
	"github.com/Twofold-One/quotes-memorizer-api/pkg/repository"
	"github.com/Twofold-One/quotes-memorizer-api/pkg/service"

	"github.com/sirupsen/logrus"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// @title           Quotes-Memorizer API
// @version         1.0
// @description     API backend for Quotes-Memorizer App.

// @contact.name   Twofold-One
// @contact.email  evangerasdev@gmail.com

// @host      https://quotes-memorizer.herokuapp.com/
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	// if err := initConfig(); err != nil {
	// 	logrus.Fatalf("error initializing configs: %s", err.Error())
	// }

	port := os.Getenv("PORT")

	// if err := godotenv.Load(); err != nil {
	// 	logrus.Fatalf("error loading env variable: %s", err.Error())
	// }

	// for local env

	// db, err := repository.NewPostgresDB(repository.Config{
	// 	Username: viper.GetString("db.username"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	Host: viper.GetString("db.host"),
	// 	Port: viper.GetString("db.port"),
	// 	DBName: viper.GetString("db.dbname"),
	// })

	db, err := repository.NewPostgresDB()
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(quotes_memorizer.Server)

	// graceful shutdown
	go func () {
		// for local env: viper.GetString("port")
		if err := srv.Run(port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("Quote Memorizer App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	logrus.Print("Quote Memorizer App shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured during server shutdown: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured during closing DB connection: %s", err.Error())
	}
}

// func initConfig() error {
// 	viper.AddConfigPath("configs")
// 	viper.SetConfigName("config")
// 	return viper.ReadInConfig()
// }