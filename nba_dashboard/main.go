package main

import (
	"net/http"

	"github.com/ivanmrnn/uadmin_sample/nba_dashboard/models"
	"github.com/ivanmrnn/uadmin_sample/nba_dashboard/views"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.Players{},
		models.Teams{},
	)
	uadmin.RootURL = "/admin/"
	uadmin.SiteName = "NBA Dashboard"

	http.HandleFunc("/", uadmin.Handler(views.MainHandler))

	uadmin.StartServer()

}
