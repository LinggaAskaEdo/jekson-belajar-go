package main

import (
	"fmt"
	"log"

	"belajar-go/src/config/query"
	"belajar-go/src/handler/rest"
	"belajar-go/src/repository"
	"belajar-go/src/service"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed load env: %v", err)
	}

	// Load DB config
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("failed load config: %v", err)
	}

	// Open PostgreSQL connection
	db, err := sqlx.Connect(cfg.Postgres.Driver, cfg.Postgres.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer func(db *sqlx.DB) {
		if cerr := db.Close(); cerr != nil {
			log.Printf("failed to close database connection: %v", cerr)
		}
	}(db)

	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime)
	db.SetConnMaxIdleTime(cfg.Postgres.ConnMaxIdleTime)

	fmt.Println("Connected to PostgreSQL")

	ql, err := query.NewLoadQuery("etc/query/user.sql")
	if err != nil {
		log.Printf("failed to load query: %v", err)
		return
	}

	userRepo := repository.InitRepository(db, ql)
	userService := service.InitService(userRepo)

	rest.InitRestHandler(userService, cfg.Server.Port)

	fmt.Println("Server running on :", cfg.Server.Port)
}
