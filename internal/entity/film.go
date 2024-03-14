package entity

import "time"

type Film struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date"`
	Rating      int       `json:"rating"`
	Actors      []Actor   `json:"actors"`
}

type FilmForm struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date"`
	Rating      int       `json:"rating"`
	Actors      []int     `json:"actors"`
}
