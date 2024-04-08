package models

import (
	"github.com/uadmin/uadmin"
)

type Team struct {
	uadmin.Model
	Name  string
	Logo string `uadmin:"image"`
	Primary string
	Secondary string
}

func (p *Team) Save() {
	uadmin.Save(p)
}