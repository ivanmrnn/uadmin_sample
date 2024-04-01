package models

import (
	"github.com/uadmin/uadmin"
)

type FolderGroup struct {
	uadmin.Model
	Group uadmin.UserGroup
	GroupID uint
	Folder Folder
	FolderID uint
	Read bool
	Add bool
	Edit bool
	Delete bool
}

func (f *FolderGroup) String() string {
	uadmin.Preload(f)
	return f.Group.GroupName
}