package entity

import "time"

type Actor struct {
	ID     int
	Name   string
	Gender string
	Birth  time.Time
}

type ActorWithFilms struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Gender string    `json:"gender"`
	Birth  time.Time `json:"birth"`
	Films  []Film    `json:"films"`
}
