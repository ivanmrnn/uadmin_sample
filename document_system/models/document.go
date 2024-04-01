package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

// Document !
type Document struct {
	uadmin.Model
	Name        string
	File        string `uadmin:"file"`
	Description string `uadmin:"html"`
	RawText     string `uadmin:"list_exclude"`
	Format      Format `uadmin:"list_exclude"`
	Folder      Folder `uadmin:"filter"`
	FolderID    uint
	CreatedDate time.Time
	Channel     Channel `uadmin:"list_exclude"`
	ChannelID   uint
	CreatedBy   string
}

func (d *Document) Save() {
	docChange := false
	newDoc := false

	// Checks whether the document record is new or existing
	if d.ID != 0 {
		// Initializes the Document model
		oldDoc := Document{}

		// Gets the ID of the old Document

	}
}
