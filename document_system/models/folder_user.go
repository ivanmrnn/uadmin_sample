package models

import (
	"github.com/uadmin/uadmin"
)

type FolderUser struct {
	uadmin.Model
	User uadmin.User
	UserID uint
	Folder Folder
	FolderID uint
	Read bool
	Add bool
	Edit bool
	Delete bool
}

func (f *FolderUser) String() string {
	uadmin.Preload(f)
	return f.User.String()
}