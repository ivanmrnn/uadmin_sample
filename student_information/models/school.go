package models

import (
	"github.com/uadmin/uadmin"
)

type School struct {
	uadmin.Model
	Name string
	Icon string `uadmin:"image"`
	Link string `uadmin:"link"`
	Site string `uadmin:"list_exclude"`
}

func (s *School) Save() {
	s.Link = s.Site
	uadmin.Save(s)
}
