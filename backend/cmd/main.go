package main

import (
	"bonus"
	"errors"
	"os"

	"bonus/pkg/handler"
	"bonus/pkg/repository"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Print("No .env file found, please set")
	}
}

// @title Database Assignment
// @version 1
// @description This is a DB assignment.
// @host http://174.129.87.192:4000
// @BasePath /
func main() {
	logrus.Print("Startup server")
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initEnv(); err != nil {
		logrus.Fatalf("error initializing env: %s", err.Error())
	}

	db, err := repository.NewPostgreDB(
		os.Getenv("DSN"),
	)

	if err != nil {
		logrus.Fatalf(err.Error())
	}

	repos := repository.NewRepository(db)
	handlers := handler.NewHandler(repos)
	/*if os.Getenv("Seed") != "" {
		err := repos.Seeder.SeedCategory()
		if err != nil {
			log.Print(err)
		}
	}*/
	srv := new(bonus.Server)

	if err := srv.Run(os.Getenv("APIPortHTTP"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initEnv() error {

	reqs := []string{
		"StaticPortHTTP",
		"APIPortHTTP",
	}

	for i := 0; i < len(reqs); i++ {
		_, exists := os.LookupEnv(reqs[i])

		if !exists {
			return errors.New(".env variables not set")
		}
	}

	return nil
}
