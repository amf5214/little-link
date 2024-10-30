package frontend

import (
	"net/http"

	g "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	html "maragu.dev/gomponents/html"
)

// A sample Gomponents handler
func GomponentsHandler(w http.ResponseWriter, r *http.Request, contents g.Node) {

	// Render the component as HTML
	err := contents.Render(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the HTML to the response
	w.Header().Set("Content-Type", "text/html")
}

func Page(title, path string, body g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []g.Node{
			html.Script(g.Attr("src", "https://cdn.tailwindcss.com")),
		},
		Body: []g.Node{
			Navbar(path),
			BodyContainer(body),
		},
	})
}

func Navbar(currentPath string) g.Node {
	return html.Nav(html.Class("navbar bg-gray-700 flex items-center space-x-4 p-4 text-white"),
		Container(
			NavbarLink("/home", "Home", currentPath == "/home"),
			NavbarLink("/home", "Manage Links", currentPath == "/manage-links"),
			NavbarLink("/home", "Settings", currentPath == "/settings"),
		),
	)
}

func NavbarLink(path, text string, active bool) g.Node {
	if active {
		return html.A(html.Class("hover:cursor-pointer mx-2 text-gray-400"), html.Href(path), g.Text(text))
	} else {
		return html.A(html.Class("hover:cursor-pointer mx-2"), html.Href(path), g.Text(text))
	}
}

func Container(children ...g.Node) g.Node {
	return html.Div(html.Class("container"), g.Group(children))
}

func BodyContainer(children ...g.Node) g.Node {
	return html.Div(html.Class("bg-gray-900 h-screen text-white flex flex-auto justify-center content-center"), g.Group(children))

}
