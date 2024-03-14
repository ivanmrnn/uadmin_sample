package models

import (
	"github.com/uadmin/uadmin"
)

type schoolCode int

func (schoolCode) BatangasStateUniversity() schoolCode {
	return 1
}

func (schoolCode) UniversityOfThePhilippines() schoolCode {
	return 2
}

func (schoolCode) UniversityOfBatangas() schoolCode {
	return 3
}

type School struct {
	uadmin.Model
	Name       string     `uadmin:"read_only"`
	SchoolCode schoolCode `uadmin:"list_exclude"`
	Icon       string     `uadmin:"image"`
	Link       string     `uadmin:"link"`
	Site       string     `uadmin:"list_exclude;read_only"`
}

func (s *School) Save() {
	s.Link = s.Site
	if s.SchoolCode == 1 {
		s.Name = "Batangas State University"
		s.Site = "https://batstateu.edu.ph/"
	} else if s.SchoolCode == 2 {
		s.Name = "University of the Philippines"
		s.Site = "https://up.edu.ph/"
	} else if s.SchoolCode == 3 {
		s.Name = "University of Batangas"
		s.Site = "https://ub.edu.ph/"
	}
	uadmin.Save(s)
}
