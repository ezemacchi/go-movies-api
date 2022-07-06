package models

import (
	"context"
	"database/sql"
	"time"
)

type GenresRepository struct {
	Db *sql.DB
}

func (m *GenresRepository) Get(id int) (*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select 
				id, genre_name, created_at, updated_ad
			from
				genres
			where
				id = $id`

	row := m.Db.QueryRowContext(ctx, query, id)

	var genre Genre

	err := row.Scan(
		&genre.ID,
		&genre.GenreName,
		&genre.CreatedAt,
		&genre.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &genre, nil
}
