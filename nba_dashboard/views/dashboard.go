package views

import (
	"net/http"

	"github.com/ivanmrnn/uadmin_sample/nba_dashboard/models"
	"github.com/uadmin/uadmin"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session) {
	type Context struct {
		User        string
		ProfilePhoto string `uadmin:"image"`
		PlayerName models.Players
		PlayerPhoto models.Players 
		Team models.Players
		PPG models.Players
		RPG models.Players
		APG models.Players
		PIE models.Players
		Birthdate models.Players
		Age models.Players
		Experience models.Players
		Height models.Players
		Weight models.Players
		Country models.Players
		Logo models.Players
		Primary models.Players
		Secondary models.Players
		Players      []models.Players 
	}

	c := Context{}
	c.User = session.User.FirstName + " " + session.User.LastName
	c.ProfilePhoto = session.User.Photo 


	players := []models.Players{}
	uadmin.All(&players)
	uadmin.Trail(uadmin.DEBUG, c.Players)


	uadmin.RenderHTML(w, r, "templates/dashboard.html", c)
}
