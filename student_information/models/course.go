package models

import (
	"github.com/uadmin/uadmin"
	"strconv"
)

type Course struct {
	uadmin.Model
	Name       string `uadmin:"required"`
	CourseCode uint `uadmin:"max:99;required;default_value: "`
	Code       string `uadmin:"hidden;list_exclude"`
}

func (c *Course) Save() {
	c.Code = strconv.FormatUint(uint64(c.CourseCode), 10)
	uadmin.Save(c)
}
