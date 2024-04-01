package main

import (
	"github.com/ivanmrnn/uadmin_sample/document_system/models"
	"github.com/uadmin/uadmin"
)

func main() {
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

	uadmin.RegisterInlines(
		models.Folder{},
		map[string]string{
			"foldergroup": "FolderID",
			"folderuser": "FolderID",
		},
	)
	uadmin.RegisterInlines(
		models.Document{},
		map[string]string{
			"documentgroup": "DocumentID",
			"documentuser": "DocumentID",
			"documentversion": "DocumentID",
		},
	)
		
	uadmin.SiteName = "Document System"
	uadmin.StartServer()
}
