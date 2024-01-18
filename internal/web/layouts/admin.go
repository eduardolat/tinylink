package layouts

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func Admin(title string, children []gomponents.Node) gomponents.Node {
	if title != "" {
		title = " - " + title
	}
	title = "TinyLink" + title

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
			html.Nav(
				html.Class("container"),
				html.Ul(
					html.Li(
						html.Img(
							html.Src("/static/images/mascot.png"),
							html.Width("40"),
						),
					),
					html.Li(
						html.Strong(
							gomponents.Text("TinyLink"),
						),
					),
				),
			),

			html.Main(
				html.Class("container"),
				gomponents.Group(children),
			),
		},
	})
}
