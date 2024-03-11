package main

import (
	"net/http"

    // Specify the username that you used inside github.com folder
    "github.com/ivanmrnn/login_system_views/views"
	"github.com/uadmin/uadmin"
)

func main() {
	// Assign RootURL value as "/admin/" and Site Name as "Login System"
	// NOTE: This code works only if database does not exist yet.
	uadmin.RootURL = "/admin/"
	uadmin.SiteName = "Login System"
	
	http.HandleFunc("/login_system/", uadmin.Handler(views.MainHandler))

	// Run the server
	uadmin.StartServer()

}
