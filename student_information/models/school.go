package models

import (
	"github.com/uadmin/uadmin"
	"strconv"
)

type School struct {
	uadmin.Model
	Name        string `uadmin:"required"`
	SchoolCode  uint   `uadmin:"max:99;required;default_value: "`
	Code        string `uadmin:"hidden;list_exclude"`
	IDFormat    IDFormat
	IDFormatID  uint
	Format      string `uadmin:"hidden;list_exclude"`
	Icon        string `uadmin:"image"`
	Website     string `uadmin:"link"`
	WebsiteLink string `uadmin:"list_exclude"`
}

func (s *School) Save() {
	uadmin.Preload(s)
	s.Format = s.IDFormat.Format
	s.Website = s.WebsiteLink
	s.Code = strconv.FormatUint(uint64(s.SchoolCode), 10)
	uadmin.Save(s)
}
