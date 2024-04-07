package models

import (
	"github.com/uadmin/uadmin"
)

type Teams struct {
	uadmin.Model
	Name  string
	Logo string `uadmin:"image"`
	PrimaryTheme string
	SecondaryTheme string
}