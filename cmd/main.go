package main

import (
	"errors"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"log"
	"vk/internal/handler"
	"vk/internal/repository"
	"vk/internal/usecase"
	"vk/pkg/server"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "quest"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("connection failed: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("connection failed: %s", err.Error())
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("somethig went wrong: %s", err.Error())
		}
	}()

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})

	if err != nil {
		log.Fatalf("somethig went wrong: %s", err.Error())

	}

	m, err := migrate.NewWithDatabaseInstance("file://database/migration", dbname, driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	}

	repo := repository.NewRepository(db)
	usecase := usecase.NewUseCase(repo)
	handlers := handler.NewHandler(usecase)

	var srv server.Server
	err = srv.Run("8080", handlers.Handle())

	if err != nil {
		log.Fatalf("connection failed: %s", err.Error())
	}
}
