package frontend

import (
	"net/http"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	components := Page("Home", "/home", InputBox("Create", "shortenUrl"))
	GomponentsHandler(w, r, components)
}

func GetCompletedPage(w http.ResponseWriter, r *http.Request, shortUrl string) {
	components := Page("Home", "/home", VBox(InputBox("Create", "shortenUrl"), html.A(html.Class("text-white p-2 "), html.Href("/"+shortUrl), g.Text("localhost:8100/"+shortUrl))))
	GomponentsHandler(w, r, components)
}
