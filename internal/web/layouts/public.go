package layouts

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func Public(title string, children []gomponents.Node) gomponents.Node {
	if title == "" {
		title = "TinyLink"
	}

	return components.HTML5(components.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []gomponents.Node{
			html.Link(
				html.Rel("stylesheet"),
				html.Href("https://cdn.jsdelivr.net/npm/@picocss/pico@1/css/pico.min.css"),
			),
		},
		Body: []gomponents.Node{
			html.Main(
				html.Class("container"),
				gomponents.Group(children),
			),
		},
	})
}
