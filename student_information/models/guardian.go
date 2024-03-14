package models

import (
	"github.com/uadmin/uadmin"
)

type Guardian struct {
	uadmin.Model
	Name    string `uadmin:"required;search"`
	Address string
	Phone   string
}
