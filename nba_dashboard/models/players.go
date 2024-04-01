package models

import (
	"github.com/uadmin/uadmin"
)

type Players struct {
	uadmin.Model
	Name  string
	PPG  int
	RPG int
	APG int
	PIE int
	Team string
	Number string
}