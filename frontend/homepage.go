package frontend

import (
	"net/http"

	g "maragu.dev/gomponents"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	components := Page("Home", "/home", g.Text("This is a sample component."))
	GomponentsHandler(w, r, components)
}
