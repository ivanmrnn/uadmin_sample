package models

import (
	"github.com/uadmin/uadmin"
)

type courseCode int

func (courseCode) Engineering() courseCode {
	return 1
}

func (courseCode) Nursing() courseCode {
	return 2
}

func (courseCode) Psychology() courseCode {
	return 3
}

func (courseCode) Education() courseCode {
	return 4
}

func (courseCode) Communications() courseCode {
	return 5
}

func (courseCode) Economics() courseCode {
	return 6
}

func (courseCode) History() courseCode {
	return 7
}

func (courseCode) PoliticalScience() courseCode {
	return 8
}

type Course struct {
	uadmin.Model
	Name       string     `uadmin:"read_only"`
	CourseCode courseCode `uadmin:"list_exclude"`
	Code       uint       `uadmin:"read_only"`
}

func (c *Course) Save() {
	switch c.CourseCode {
	case 1:
		c.Name = "Engineering"
		c.Code = 1
	case 2:
		c.Name = "Nursing"
		c.Code = 2
	case 3:
		c.Name = "Psychology"
		c.Code = 3
	case 4:
		c.Name = "Education"
		c.Code = 4
	case 5:
		c.Name = "Communication"
		c.Code = 5
	case 6:
		c.Name = "Economics"
		c.Code = 6
	case 7:
		c.Name = "History"
		c.Code = 7
	case 8:
		c.Name = "Political Science"
		c.Code = 8
	}
	uadmin.Save(c)
}
