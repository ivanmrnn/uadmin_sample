package models

import (
	"github.com/uadmin/uadmin"
)

type Teams struct {
	uadmin.Model
	Name  string
	PPG  int
	RPG int
	APG int
	Founded string
	City string
}