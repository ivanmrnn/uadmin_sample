package models

import (
	"github.com/uadmin/uadmin"
)

type Team struct {
	uadmin.Model
	Name string
	Apg  int
}

func (t *Team) Save() {
	players := []Player{}
	uadmin.All(&players)
	avg := 0
	for i := range players {
		if players[i].TeamName == 
		avg +=  players[i].Apg
	}
	t.Apg = avg
	uadmin.Save(t)
}
