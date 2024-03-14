package entity

import "time"

type Actor struct {
	ID     int
	Name   string
	Gender string
	Birth  time.Time
}
