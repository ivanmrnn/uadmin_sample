package main

import (
	"net/http"

	"github.com/ivanmrnn/to-do/api"
	"github.com/ivanmrnn/to-do/models"
	"github.com/ivanmrnn/to-do/views"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.Todo{},
		models.Category{},
		models.Friend{},
		models.Item{},
	)
	uadmin.RegisterInlines(models.Category{}, map[string]string{
		"Todo": "CategoryID",
	})
	uadmin.RegisterInlines(models.Friend{}, map[string]string{
		"Todo": "FriendID",
	})
	uadmin.RegisterInlines(models.Item{}, map[string]string{
		"Todo": "ItemID",
	})

	// API Handler
	http.HandleFunc("/api/", uadmin.Handler(api.Handler))
	http.HandleFunc("/page/", uadmin.Handler(views.PageHandler))

	uadmin.StartSecureServer("pub.pem", "priv.pem")
}
