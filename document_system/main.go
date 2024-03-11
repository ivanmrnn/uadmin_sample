package main

import (
	"github.com/ivanmrnn/uadmin_sample/document_system/models"
	"github.com/uadmin/uadmin"
)

func main() {
	// Assign Site Name value as "Document System"
	// NOTE: This code works only if database does not exist yet.
	uadmin.SiteName = "Document System"

	uadmin.Register(
		models.Folder{},
		models.FolderGroup{},
		models.FolderUser{},
		models.Channel{},
		models.Document{},
		models.DocumentGroup{},
		models.DocumentUser{},
		models.DocumentVersion{},
	)
	// Register FolderGroup and FolderUser to Folder model
	uadmin.RegisterInlines(
		models.Folder{},
		map[string]string{
			"foldergroup": "FolderID",
			"folderuser":  "FolderID",
		},
	)

	// Register DocumentVersion, DocumentGroup, and DocumentUser to Document
	// model
	uadmin.RegisterInlines(
		models.Document{},
		map[string]string{
			"documentgroup":   "DocumentID",
			"documentuser":    "DocumentID",
			"documentversion": "DocumentID",
		},
	)

	// Activates a uAdmin server
	uadmin.StartServer()
}
