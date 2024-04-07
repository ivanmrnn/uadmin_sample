package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Players struct {
	uadmin.Model
	Name  string
	Photo string `uadmin:"image"`
	Team Teams
	PPG  int
	RPG int
	APG int
	PIE int
	Birthdate time.Time
	Age int
	Experience int
	Height int
	Weight int
	Country string
}