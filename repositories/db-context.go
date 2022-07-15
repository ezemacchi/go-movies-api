package repositories

import (
	"context"
	"database/sql"
	"time"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type DbContext struct {
	Db *sql.DB
}

func (app *application) openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
