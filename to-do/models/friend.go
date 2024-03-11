package models

import (
	"github.com/uadmin/uadmin"
)

// Friend Model !
type Friend struct {
	uadmin.Model
	Name        string `uadmin:"required"`
	Email       string `uadmin:"email"`
	Password    string `uadmin:"password;list_exclude;encrypt"`
	Nationality Nationality
	Invite      string `uadmin:"link"`
}

func (f *Friend) Save() {
	f.Invite = "https://www.google.com/"
	uadmin.Save(f)
}

type Nationality int

// Chinese !
func (Nationality) Chinese() Nationality {
	return 1
}

// Filipino !
func (Nationality) Filipino() Nationality {
	return 2
}

// Others !
func (Nationality) Others() Nationality {
	return 3
}
