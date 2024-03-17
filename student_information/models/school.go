package models

import (
	"github.com/uadmin/uadmin"
)

type School struct {
	uadmin.Model
	Name        string
	SchoolCode  uint
	Format      string
	Icon        string `uadmin:"image"`
	Website     string `uadmin:"link"`
	WebsiteLink string `uadmin:"list_exclude"`
}

func (s *School) Save() {
	s.Website = s.WebsiteLink
	uadmin.Save(s)
}
