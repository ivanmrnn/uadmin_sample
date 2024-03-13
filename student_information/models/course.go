package models

import (
	"github.com/uadmin/uadmin"
)

type Course struct {
	uadmin.Model
	Name    string 
}