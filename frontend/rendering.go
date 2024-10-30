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
		Body: []g.Node{
			Navbar(path),
			Container(body),
		},
	})
}

func Navbar(currentPath string) g.Node {
	return html.Nav(html.Class("navbar"),
		Container(
			NavbarLink("/home", "Home", currentPath == "/home"),
		),
	)
}

func NavbarLink(path, text string, active bool) g.Node {
	return html.A(html.Href(path), g.Text(text),
		c.Classes{
			"active": active,
		},
	)
}

func Container(children ...g.Node) g.Node {
	return html.Div(html.Class("container"), g.Group(children))
}