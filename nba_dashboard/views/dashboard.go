package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session) {
	type Context struct {
		User        string
		OTPRequired bool
	}

	c := Context{}
	c.User = session.User.FirstName + " " + session.User.LastName

	if session.User.OTPRequired {
		c.OTPRequired = true
	}

	uadmin.RenderHTML(w, r, "templates/dashboard.html", c)
	return
}
