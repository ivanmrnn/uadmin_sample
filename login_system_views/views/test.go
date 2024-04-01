package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

// HomeHandler handles the home page.
func TestHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session) {
	// Initialize the fields that we need in the custom struct.
	type Context struct {
		Name        string	
	}

	// Call the custom struct and assign the full name in the User field under the context object.
	c := Context{}
	c.Name = session.User.FirstName + " " + session.User.LastName

	// Check if the user has OTPRequired enabled in the database.


	// Render the home filepath and pass the context data object to the HTML file.
	uadmin.RenderHTML(w, r, "templates/test.html", c)
	return
}
