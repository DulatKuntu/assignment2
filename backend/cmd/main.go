package main

import (
	"bonus"
	"errors"
	"os"

	"bonus/model"
	"bonus/pkg/handler"
	"bonus/pkg/repository"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Print("No .env file found, please set")
	}
}

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
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	err = migrateGorm(gormDB)
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	repos := repository.NewRepository(db, gormDB)
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

func migrateGorm(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.Country{},
		&model.DiseaseType{},
		&model.Disease{},
		&model.Country{},
		&model.Discover{},
		&model.Users{},
		&model.PublicServant{},
		&model.Doctor{},
		&model.Specialize{},
		&model.Record{},
	)
	if err != nil {
		return err
	}
	return nil
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
