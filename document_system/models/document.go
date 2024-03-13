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
	// Initialized variables
	docChange := false
	newDoc := false
	// Checks whether the document record is new or existing
	if d.ID != 0 {
		// Initializes the Document model
		oldDoc := Document{}

		// Gets the ID of the old Document
		uadmin.Get(&oldDoc, "id = ?", d.ID)

		// Checks if the file is changed or updated
		if d.File != oldDoc.File {
			docChange = true
		}
	} else {
		// New document record
		docChange = true
		newDoc = true
	}

	// Save the Document
	uadmin.Save(d)

	// Checks whether the document record has changed
	if docChange {
		// Sets the document value to the DocumentVersion fields
		ver := DocumentVersion{}
		ver.Date = time.Now()
		ver.DocumentID = d.ID
		ver.File = d.File
		ver.Format = d.Format

		// Counts the version number based on the DocumentID and increment it
		// by 1
		ver.Number = uadmin.Count([]DocumentVersion{}, "document_id = ?", d.ID) + 1

		// Save the document version
		uadmin.Save(&ver)

		if newDoc {
			// Initializes the User model
			user := uadmin.User{}

			// Gets the username of the user to display in CreatedBy
			uadmin.Get(&user, "username = ?", d.CreatedBy)

			// Sets values to the DocumentUser model fields
			creator := DocumentUser{
				UserID:     user.ID,
				DocumentID: d.ID,
				Read:       true,
				Edit:       true,
				Add:        true,
				Delete:     true,
			}

			// Save the document user
			uadmin.Save(&creator)
		}
	}
}
