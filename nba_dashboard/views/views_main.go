package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// MainHandler is the main handler for the login system.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called "/login_system/"
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/nba_dashboard")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
	session := uadmin.IsAuthenticated(r)
	
	if session == nil {
		LoginHandler(w, r)
		return
	}

	if r.URL.Path == "" {
		DashboardHandler(w, r, session)
		return

	} else if r.URL.Path == "/logout" {
		/* If the request URL Path is /logout after the /login_system/, it will proceed to this part.
		   e.g. localhost:8080/login_system/logout */

		// LogoutHandler handles the logout process for the user.
		LogoutHandler(w, r, session)
		return
	}
	return
}