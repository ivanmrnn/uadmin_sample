package models

import (
	"github.com/uadmin/uadmin"
)

type Player struct {
	uadmin.Model
	Name string
	Apg  int

	Team   Team
	TeamID uint
	TeamName string
	
}

func (t *Player) Save() {
	uadmin.Preload(t)
	t.TeamName = t.Team.Name
	uadmin.Save(t)
}
