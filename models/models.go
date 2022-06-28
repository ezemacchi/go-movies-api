package models

import "time"

type Movie struct {
	Id          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"descrition"`
	Year        int          `json:"year"`
	ReleaseDate time.Time    `json:"release_date"`
	Runtime     int          `json:"runtime"`
	Rating      int          `json:"rating"`
	MPAARating  string       `json:"mpaa_rating"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"update_at"`
	MovieGenre  []MovieGenre `json:"-"`
}

type Genre struct {
	Id        int       `json:"int"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

type MovieGenre struct {
	Id           int       `json:"id"`
	MovieId      int       `json:"movie_id"`
	MovieGenreId int       `json:"movie_genre_id"`
	Genre        Genre     `json:"genre"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"update_at"`
}
