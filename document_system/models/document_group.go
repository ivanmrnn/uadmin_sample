package models

import (
	"github.com/uadmin/uadmin"
)

type DocumentGroup struct {
	uadmin.Model
	Group uadmin.UserGroup
	GroupID uint
	Document Document
	DocumentID uint
	Read bool
	Add bool
	Edit bool
	Delete bool
}

// DocumentGroup function that returns string value
func (d *DocumentGroup) String() string {

	// Gives access to the fields in another model
	uadmin.Preload(d)

	// Returns the full name from the User model
	return d.Group.GroupName
}