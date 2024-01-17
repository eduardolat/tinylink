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
		Head:     []gomponents.Node{},
		Body: []gomponents.Node{
			html.Main(children...),
		},
	})
}
