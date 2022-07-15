package main

import (
	"flag"
	"fmt"
	"log"
	"movies-api/models"
	"movies-api/repositories"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type Application struct {
	config    config
	logger    *log.Logger
	models    models.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://admin:admin@localhost/movies_db?sslmode=disable", "Postgres connection string")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModels(db),
		dbContext: repositories.DbContext.,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", cfg.port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
