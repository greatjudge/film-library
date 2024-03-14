package entity

import "time"

type Film struct {
	ID          int
	Title       string
	Description string
	ReleaseDate time.Time
	Rating      int
	Actors      []Actor
}
