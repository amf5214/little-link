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
			// html.Link(g.Attr("href", "./output.css"), g.Attr("rel", "stylesheet")),
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

func VBox(children ...g.Node) g.Node {
	return html.Div(html.Class("flex flex-col justify-center content-center items-center flex-auto "), g.Group(children))
}

func HBox(classString string, children ...g.Node) g.Node {
	return html.Div(html.Class("flex flex-row justify-center content-center items-center flex-auto "+classString), g.Group(children))
}

func InputBox(text string, buttonFunction string, children ...g.Node) g.Node {
	return html.Form(g.Attr("action", "/tinylink"), g.Attr("method", "POST"), HBox("max-h-[30vh] max-w-[30vw] p-3 bg-gray-800", html.Input(html.Class("mr-1 h-16 w-64 p-1 bg-white text-black"), g.Attr("name", "urlpath")), html.Input(html.Class("ml-1 h-16 w-32 p-1 bg-white text-black"), g.Text("Create"), g.Attr("type", "submit"))))

}

func BodyContainer(children ...g.Node) g.Node {
	return html.Div(html.Class("bg-gray-900 h-screen text-white flex flex-auto justify-center content-center items-center"), g.Group(children))

}
